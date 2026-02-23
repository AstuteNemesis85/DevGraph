package analysis

// complexity.go — Big-O time and space complexity inference
//
// Public API (signature unchanged for worker.go compatibility)
// ──────────────────────────────────────────────────────────
//   inferComplexity(code string) (timeComplexity, spaceComplexity string)
//
// Both return values are Big-O strings stored directly in the CodeAnalysis
// row (TimeComplexity / SpaceComplexity columns of the database).
//
// Implementation
// ──────────────
// inferComplexity delegates to analyzeCode (defined in detect.go) to obtain
// the shared codeFeatures struct, then applies two independent decision trees:
//   classifyTime  — derives time  complexity from algorithmic patterns + loop depth
//   classifySpace — derives space complexity from data-structure usage + recursion depth

// inferComplexity is the package-level entry point called by worker.go.
// It scans the source once through analyzeCode and applies both decision trees.
func inferComplexity(code string) (timeComplexity, spaceComplexity string) {
	f := analyzeCode(code)
	return classifyTime(f), classifySpace(f)
}

// ─────────────────────────────────────────────────────────────────────────────
// Time-complexity decision tree
// ─────────────────────────────────────────────────────────────────────────────
//
// Rules are ordered from highest to lowest complexity so that the worst-case
// class wins whenever multiple patterns co-exist.  This mirrors how humans
// reason about upper bounds.

func classifyTime(f codeFeatures) string {
	switch {

	// ── O(n^2): 2-D DP table (nested loops + dp[i][j]) ───────────────────
	case f.hasDPTable && f.maxLoopDepth >= 2:
		return "O(n^2)"

	// ── O(n log n): divide-and-conquer with a sort call
	//    (merge sort, Tim sort on sub-problems, etc.) ──────────────────────
	case f.hasDivideConquer && f.hasSorting:
		return "O(n log n)"

	// ── O(n log n): pure divide-and-conquer (merge sort, segment trees) ──
	case f.hasDivideConquer:
		return "O(n log n)"

	// ── O(n log n): linear scan containing a binary search ───────────────
	case f.hasBinarySearch && f.maxLoopDepth >= 1:
		return "O(n log n)"

	// ── O(2^n): bare (unmemoised) recursion — exponential growth ─────────
	//   Requires that none of the sub-linear or polynomial optimisations
	//   (memoisation, D&C, binary search) are present.
	case f.hasRecursion && !f.hasDPMemo && !f.hasBinarySearch && !f.hasDivideConquer:
		return "O(2^n)"

	// ── O(log n): isolated binary search ─────────────────────────────────
	case f.hasBinarySearch:
		return "O(log n)"

	// ── O(n^2): memoised recursion (conservative upper-bound)
	//    Many DP problems are O(n) or O(n·k) but without knowing state
	//    dimensions we default to the common quadratic case.
	case f.hasDPMemo:
		return "O(n^2)"

	// ── O(n^2): structurally nested loops ────────────────────────────────
	case f.maxLoopDepth >= 2:
		return "O(n^2)"

	// ── O(n log n): single loop containing an in-loop sort call ──────────
	case f.maxLoopDepth == 1 && f.hasSorting:
		return "O(n log n)"

	// ── O(n): single loop or linear DFS/BFS ──────────────────────────────
	case f.maxLoopDepth == 1 || f.hasDFSBFS:
		return "O(n)"

	// ── O(1): no loops, no recursion, no traversal ───────────────────────
	default:
		return "O(1)"
	}
}

// ─────────────────────────────────────────────────────────────────────────────
// Space-complexity decision tree
// ─────────────────────────────────────────────────────────────────────────────
//
// Space complexity is driven by:
//   a) Explicit data structures allocated proportional to input size.
//   b) Implicit recursion call-stack depth.
//
// The tree considers both dimensions.

func classifySpace(f codeFeatures) string {
	switch {

	// ── O(n^2): 2-D array / vector-of-vectors ────────────────────────────
	//   Dominates because the structure alone requires n² cells.
	case f.uses2DArray:
		return "O(n^2)"

	// ── O(n): memoisation map holds one entry per unique sub-problem ──────
	case (f.hasDPMemo || f.hasDPTable) && f.usesMap:
		return "O(n)"

	// ── O(log n): balanced recursive call stack (binary search, D&C) ─────
	//   The recursion stack depth is O(log n) when the input is halved each
	//   level, so no heap allocation is needed beyond the stack frames.
	case f.hasRecursion && f.hasDivideConquer:
		return "O(log n)"

	// ── O(n): linear recursion stack (e.g. DFS on a path-shaped graph) ───
	case f.hasRecursion:
		return "O(n)"

	// ── O(n): heap-allocated linear structures ────────────────────────────
	case f.usesMap || f.usesVector || f.usesStack || f.usesQueue:
		return "O(n)"

	// ── O(1): no heap allocation, no recursion ────────────────────────────
	default:
		return "O(1)"
	}
}
