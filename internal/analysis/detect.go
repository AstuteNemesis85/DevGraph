package analysis

// detect.go — Structural pattern detector
//
// Architecture
// ────────────
// Rather than counting raw keyword occurrences we run a lightweight
// single-pass structural scanner that:
//
//  1. Strips comments and string literals (prevents false positives).
//  2. Tokenises the cleaned source to track brace scopes, building a
//     scope stack that knows whether each open-brace belongs to a loop.
//  3. Derives codeFeatures from the structural scan plus targeted
//     pre-compiled regular expressions for specific sub-patterns.
//
// codeFeatures is the single shared intermediate representation; both
// detectPatterns (public) and inferComplexity (public) call analyzeCode
// internally so the source is scanned with the same logic.
//
// Public API (unchanged for worker.go compatibility)
// ──────────────────────────────────────────────────
//   detectPatterns(code string) []string

import (
	"regexp"
	"strings"
	"unicode"
)

// ─────────────────────────────────────────────────────────────────────────────
// Shared intermediate representation
// ─────────────────────────────────────────────────────────────────────────────

// codeFeatures captures every structural signal extracted from a source file.
// It is intentionally a plain struct (no mutexes, no allocations beyond slices)
// so that the analyzeCode pipeline stays allocation-light and O(n) in source
// length.
type codeFeatures struct {
	// ── Loop structure ────────────────────────────────────────────────────
	maxLoopDepth  int // deepest simultaneous loop nesting (true structural)
	numLoopBlocks int // total distinct loop-opening braces encountered

	// ── Algorithm patterns ────────────────────────────────────────────────
	hasRecursion     bool
	hasBinarySearch  bool
	hasDivideConquer bool
	hasDPMemo        bool // recursion + hash-map  → memoisation DP
	hasDPTable       bool // dp[] array accessed inside ≥2 nested loops → tabulation DP
	hasDFSBFS        bool
	hasSorting       bool
	hasHashing       bool
	hasEarlyBreak    bool

	// ── Data-structure presence (drives space complexity in complexity.go) ─
	usesVector  bool
	usesMap     bool
	uses2DArray bool // 2-D array or vector-of-vectors
	usesStack   bool // explicit stack<> / Stack usage
	usesQueue   bool // queue<> / deque<> / Queue usage

	// ── Derived helpers ───────────────────────────────────────────────────
	functionNames []string // user-defined function identifiers found in source
}

// ─────────────────────────────────────────────────────────────────────────────
// Public API (signature unchanged)
// ─────────────────────────────────────────────────────────────────────────────

// detectPatterns analyses source code and returns a deduplicated slice of
// human-readable pattern names.  These are stored as AlgorithmPattern rows;
// the names must remain stable between releases.
func detectPatterns(code string) []string {
	f := analyzeCode(code)
	return buildPatternList(f)
}

// ─────────────────────────────────────────────────────────────────────────────
// Core analysis pipeline
// ─────────────────────────────────────────────────────────────────────────────

// analyzeCode is the main entry-point for the analysis pipeline. It is called by
// BOTH detectPatterns and inferComplexity so the two public functions stay
// decoupled while sharing the same structural signals.
func analyzeCode(code string) codeFeatures {
	clean := stripCommentsAndStrings(code)

	var f codeFeatures
	f.functionNames = extractFunctionNames(clean)
	f.maxLoopDepth, f.numLoopBlocks = analyzeLoopStructure(clean)

	// ── Data structures ───────────────────────────────────────────────────
	f.usesVector  = matchesAny(clean, rxVector)
	f.usesMap     = matchesAny(clean, rxMap)
	f.uses2DArray = matchesAny(clean, rxArray2D)
	f.usesStack   = matchesAny(clean, rxStackDS)
	f.usesQueue   = matchesAny(clean, rxQueueDS)

	// ── Algorithm patterns ────────────────────────────────────────────────
	f.hasSorting      = matchesAny(clean, rxSort)
	f.hasHashing      = f.usesMap || matchesAny(clean, rxHashSet)
	f.hasBinarySearch = detectBinarySearch(clean)
	f.hasRecursion    = detectRecursion(clean, f.functionNames)
	f.hasDivideConquer = detectDivideConquer(clean, f.hasRecursion)
	f.hasDFSBFS       = detectDFSBFS(clean, f.hasRecursion, f.usesStack, f.usesQueue)
	f.hasDPMemo       = f.hasRecursion && f.usesMap
	f.hasDPTable      = detectDPTable(clean, f.maxLoopDepth)
	f.hasEarlyBreak   = matchesAny(clean, rxEarlyBreak)

	return f
}

// ─────────────────────────────────────────────────────────────────────────────
// Pre-compiled regular expressions
// ─────────────────────────────────────────────────────────────────────────────
// All regexes are compiled once at package init time (via MustCompile).
// Grouping them here makes it easy to extend to new languages: add a new
// entry to the relevant slice.

var (
	// Data structures — C++, Java, Go, Python variants
	rxVector  = rxList(`\bvector\s*<`, `\[\]`, `ArrayList\b`, `\[\s*\]`)
	rxMap     = rxList(`\bunordered_map\s*<`, `\bmap\s*<`, `\bHashMap\b`, `\bdict\b`, `\bmap\[`)
	rxArray2D = rxList(`\[\s*\w+\s*\]\s*\[`, `vector\s*<\s*vector`, `\[\]\[\]`)
	rxStackDS = rxList(`\bstack\s*<`, `\bStack\b`)
	rxQueueDS = rxList(`\bqueue\s*<`, `\bdeque\s*<`, `\bQueue\b`, `ArrayDeque\b`)

	// Algorithm markers
	rxSort      = rxList(`\bsort\s*\(`, `\.sort\s*\(`, `Collections\.sort\b`, `Arrays\.sort\b`)
	rxHashSet   = rxList(`\bunordered_set\s*<`, `\bset\s*<`, `\bHashSet\b`)
	rxEarlyBreak = rxList(`\bbreak\b`, `\breturn\b`) // presence of break is the signal

	// Function declarations — C++ / Java / Go / Python
	// Captures the function name as group 1.
	rxFuncDecl = regexp.MustCompile(`(?m)(?:^|\n)\s*(?:[\w:*&<>\[\]]+\s+)+(\w+)\s*\(`)

	// Binary-search structural landmarks
	rxBSMid  = regexp.MustCompile(`\bmid\s*=\s*[^;\n]*(?:lo|hi|low|high|left|right|l|r)\b`)
	rxBSMove = regexp.MustCompile(`(?:lo|low|left|l)\s*=\s*mid[\s]*[+\-][\s]*1|(?:hi|high|right|r)\s*=\s*mid[\s]*[+\-][\s]*1`)

	// Divide-and-conquer: recursive call that halves the input range
	rxDnC = regexp.MustCompile(`(?:n|size|len|length|count|mid|hi|high|right)\s*/\s*2`)

	// Graph traversal markers
	rxVisited = regexp.MustCompile(`\bvisited\b`)
	rxGraph   = regexp.MustCompile(`\b(?:adj|graph|neighbors|neighbours|edges|children|nodes)\b`)

	// DP table access: dp[...][...] or dp[i] / dp[i+1]
	rxDPAccess = regexp.MustCompile(`\bdp\s*\[`)
)

// rxList is a compile-time constructor that turns a variadic list of pattern
// strings into a []*regexp.Regexp.  Panics if any pattern is invalid (caught
// at startup, not at runtime).
func rxList(patterns ...string) []*regexp.Regexp {
	out := make([]*regexp.Regexp, len(patterns))
	for i, p := range patterns {
		out[i] = regexp.MustCompile(p)
	}
	return out
}

// matchesAny returns true if the cleaned source matches ANY of the supplied
// pre-compiled regexes.
func matchesAny(code string, rxs []*regexp.Regexp) bool {
	for _, rx := range rxs {
		if rx.MatchString(code) {
			return true
		}
	}
	return false
}

// ─────────────────────────────────────────────────────────────────────────────
// Comment and string-literal stripping
// ─────────────────────────────────────────────────────────────────────────────

// stripCommentsAndStrings removes:
//   - C/C++/Java // line comments
//   - C/C++/Java /* … */ block comments
//   - Double-quoted string literals  "…"
//   - Single-quoted char literals    '…'
//
// Removed content is replaced with spaces so that character offsets and
// line numbers are preserved (important for future line-level analysis).
// The function is a single O(n) pass with no regexp overhead.
func stripCommentsAndStrings(src string) string {
	var b strings.Builder
	b.Grow(len(src))
	i, n := 0, len(src)

	for i < n {
		ch := src[i]

		// Block comment /* … */
		if i+1 < n && ch == '/' && src[i+1] == '*' {
			b.WriteByte(' ')
			i += 2
			for i < n {
				if i+1 < n && src[i] == '*' && src[i+1] == '/' {
					b.WriteByte(' ')
					i += 2
					break
				}
				if src[i] == '\n' {
					b.WriteByte('\n') // preserve line structure
				} else {
					b.WriteByte(' ')
				}
				i++
			}
			continue
		}

		// Line comment // …
		if i+1 < n && ch == '/' && src[i+1] == '/' {
			b.WriteByte(' ')
			i += 2
			for i < n && src[i] != '\n' {
				b.WriteByte(' ')
				i++
			}
			continue
		}

		// String literal "…"
		if ch == '"' {
			b.WriteByte(' ')
			i++
			for i < n && src[i] != '"' {
				if src[i] == '\\' {
					b.WriteByte(' ')
					i++ // skip escaped char
				}
				b.WriteByte(' ')
				i++
			}
			if i < n {
				b.WriteByte(' ')
				i++ // closing "
			}
			continue
		}

		// Char literal '…'
		if ch == '\'' {
			b.WriteByte(' ')
			i++
			for i < n && src[i] != '\'' {
				if src[i] == '\\' {
					b.WriteByte(' ')
					i++
				}
				b.WriteByte(' ')
				i++
			}
			if i < n {
				b.WriteByte(' ')
				i++
			}
			continue
		}

		b.WriteByte(ch)
		i++
	}
	return b.String()
}

// ─────────────────────────────────────────────────────────────────────────────
// Function-name extraction
// ─────────────────────────────────────────────────────────────────────────────

// extractFunctionNames returns deduplicated user-defined function identifiers.
// It uses a single regex over the cleaned source and filters out known
// language keywords so that constructs like "if (" don't pollute the list.
func extractFunctionNames(clean string) []string {
	matches := rxFuncDecl.FindAllStringSubmatch(clean, -1)
	seen := make(map[string]struct{}, len(matches))
	names := make([]string, 0, len(matches))
	for _, m := range matches {
		if len(m) >= 2 {
			name := m[1]
			if isLangKeyword(name) {
				continue
			}
			if _, dup := seen[name]; !dup {
				seen[name] = struct{}{}
				names = append(names, name)
			}
		}
	}
	return names
}

// isLangKeyword reports whether s is a C++/Java/Go keyword that the function
// declaration regex might mistakenly capture (e.g. "if (", "while (").
var langKeywords = map[string]struct{}{
	"if": {}, "else": {}, "for": {}, "while": {}, "do": {}, "switch": {},
	"case": {}, "return": {}, "break": {}, "continue": {}, "class": {},
	"struct": {}, "namespace": {}, "template": {}, "typename": {},
	"int": {}, "long": {}, "double": {}, "float": {}, "bool": {},
	"void": {}, "auto": {}, "const": {}, "static": {}, "new": {}, "delete": {},
	"try": {}, "catch": {}, "throw": {}, "func": {}, "def": {},
}

func isLangKeyword(s string) bool {
	_, ok := langKeywords[s]
	return ok
}

// ─────────────────────────────────────────────────────────────────────────────
// Structural loop-nesting analysis
// ─────────────────────────────────────────────────────────────────────────────

// analyzeLoopStructure performs a single character-level scan that tracks a
// scope stack. Each entry on the stack records whether the corresponding
// open-brace was introduced by a loop keyword (for / while / do).
//
// Returns:
//   maxDepth   — the highest simultaneous loop-nesting level reached
//   numBlocks  — total number of loop-opening braces (≥2 with maxDepth==1
//                means sequential, non-nested loops)
func analyzeLoopStructure(clean string) (maxDepth, numBlocks int) {
	type frame struct{ isLoop bool }

	stack     := make([]frame, 0, 32)
	loopDepth := 0
	prevWord  := ""
	i, n      := 0, len(clean)

	for i < n {
		ch := clean[i]

		switch {
		case ch == '{':
			_, isLoop := loopKeywords[prevWord]
			stack = append(stack, frame{isLoop: isLoop})
			if isLoop {
				loopDepth++
				numBlocks++
				if loopDepth > maxDepth {
					maxDepth = loopDepth
				}
			}
			prevWord = ""
			i++

		case ch == '}':
			if len(stack) > 0 {
				if stack[len(stack)-1].isLoop {
					loopDepth--
				}
				stack = stack[:len(stack)-1]
			}
			prevWord = ""
			i++

		case isIdentStart(ch):
			j := i
			for j < n && isIdentContinue(clean[j]) {
				j++
			}
			prevWord = clean[i:j]
			i = j

		default:
			// Any non-identifier, non-brace character resets the "previous word"
			// only when it's not whitespace — whitespace is allowed between a loop
			// keyword and its opening brace (e.g.  "for (...)\n{").
			if !unicode.IsSpace(rune(ch)) {
				prevWord = ""
			}
			i++
		}
	}
	return
}

// loopKeywords is the set of tokens that open a new loop scope.
var loopKeywords = map[string]struct{}{
	"for": {}, "while": {}, "do": {},
}

func isIdentStart(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func isIdentContinue(c byte) bool {
	return isIdentStart(c) || (c >= '0' && c <= '9')
}

// ─────────────────────────────────────────────────────────────────────────────
// Specific algorithm-pattern detectors
// ─────────────────────────────────────────────────────────────────────────────

// detectRecursion reports true if any extracted function name appears more
// than once in the cleaned source — once for its declaration and once (or more)
// as a self-call.
func detectRecursion(clean string, funcNames []string) bool {
	for _, name := range funcNames {
		rx := regexp.MustCompile(`\b` + regexp.QuoteMeta(name) + `\s*\(`)
		if len(rx.FindAllString(clean, -1)) >= 2 {
			return true
		}
	}
	return false
}

// detectBinarySearch looks for the structural landmark pair that uniquely
// identifies binary-search logic regardless of identifier naming:
//   1. A midpoint calculation involving lo/hi style variables.
//   2. A range-halving move:  lo = mid+1  or  hi = mid-1.
func detectBinarySearch(clean string) bool {
	return rxBSMid.MatchString(clean) && rxBSMove.MatchString(clean)
}

// detectDivideConquer identifies divide-and-conquer by requiring BOTH:
//   • Recursion (a function calls itself).
//   • An explicit halving of the problem size (n/2, size/2, etc.).
func detectDivideConquer(clean string, hasRecursion bool) bool {
	return hasRecursion && rxDnC.MatchString(clean)
}

// detectDFSBFS identifies graph traversal by requiring:
//   • A "visited" guard (almost universal in graph traversal).
//   • At least one graph-vocabulary identifier (adj, graph, neighbors …).
//   • Either recursion (DFS) OR an explicit queue/stack (BFS or iterative DFS).
func detectDFSBFS(clean string, hasRecursion, usesStack, usesQueue bool) bool {
	if !rxVisited.MatchString(clean) || !rxGraph.MatchString(clean) {
		return false
	}
	return hasRecursion || usesStack || usesQueue
}

// detectDPTable identifies bottom-up dynamic programming by requiring:
//   • A dp[] array access pattern anywhere in the source.
//   • A maximum loop-nesting depth ≥ 2, indicating nested iteration over
//     state dimensions.
func detectDPTable(clean string, maxLoopDepth int) bool {
	return maxLoopDepth >= 2 && rxDPAccess.MatchString(clean)
}

// ─────────────────────────────────────────────────────────────────────────────
// Pattern list builder
// ─────────────────────────────────────────────────────────────────────────────

// buildPatternList converts codeFeatures into the stable human-readable names
// stored in the algorithm_patterns table.  The helper closure keeps the
// conditional-append logic readable without allocating per call.
func buildPatternList(f codeFeatures) []string {
	p := make([]string, 0, 8)

	maybeAdd := func(cond bool, name string) {
		if cond {
			p = append(p, name)
		}
	}

	// Loop topology — mutually exclusive labels ordered from specific to general
	maybeAdd(f.maxLoopDepth >= 2, "Nested Loop")
	maybeAdd(f.maxLoopDepth == 1 && f.numLoopBlocks >= 2, "Sequential Loops")
	maybeAdd(f.maxLoopDepth == 1 && f.numLoopBlocks == 1, "Loop")

	// Named algorithmic patterns
	maybeAdd(f.hasRecursion,     "Recursion")
	maybeAdd(f.hasBinarySearch,  "Binary Search")
	maybeAdd(f.hasSorting,       "Sorting")
	maybeAdd(f.hasHashing,       "Hashing")
	maybeAdd(f.hasDFSBFS,        "DFS/BFS")
	maybeAdd(f.hasDivideConquer, "Divide and Conquer")
	maybeAdd(f.hasDPMemo,        "Dynamic Programming (Memoization)")
	maybeAdd(f.hasDPTable,       "Dynamic Programming (Tabulation)")
	maybeAdd(f.hasEarlyBreak,    "Early Break Optimization")

	return p
}
