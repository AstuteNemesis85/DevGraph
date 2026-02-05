import Editor from '@monaco-editor/react';
import { AlertCircle, CheckCircle, Clock, Code2, Loader2, Play, X, Zap, Sparkles } from 'lucide-react';
import { useEffect, useState } from 'react';
import toast from 'react-hot-toast';
import { useNavigate } from 'react-router-dom';
import Navbar from '../components/Navbar';
import { codeAPI } from '../services/api';

const LANGUAGES = [
  { value: 'python', label: 'Python', color: 'from-blue-500 to-cyan-500' },
  { value: 'javascript', label: 'JavaScript', color: 'from-yellow-500 to-orange-500' },
  { value: 'java', label: 'Java', color: 'from-red-500 to-orange-600' },
  { value: 'cpp', label: 'C++', color: 'from-purple-500 to-pink-500' },
  { value: 'go', label: 'Go', color: 'from-cyan-500 to-blue-500' },
];

export default function Dashboard() {
  const [code, setCode] = useState('// Write your code here...\n\n');
  const [language, setLanguage] = useState('python');
  const [submitting, setSubmitting] = useState(false);
  const [submissions, setSubmissions] = useState([]);
  const [selectedSubmission, setSelectedSubmission] = useState(null);
  const [analysisData, setAnalysisData] = useState(null);
  const [loadingAnalysis, setLoadingAnalysis] = useState(false);
  const navigate = useNavigate();

  useEffect(() => {
    loadSubmissions();
  }, []);

  const loadSubmissions = async () => {
    try {
      const response = await codeAPI.getSubmissions();
      const formattedSubmissions = response.data.map(sub => ({
        id: sub.id,
        language: sub.language,
        timestamp: sub.created_at,
        status: 'completed',
      }));
      setSubmissions(formattedSubmissions);
    } catch (error) {
      console.error('Failed to load submissions:', error);
    }
  };

  const viewAnalysis = async (submissionId) => {
    setSelectedSubmission(submissionId);
    setLoadingAnalysis(true);
    try {
      const response = await codeAPI.getAnalysis(submissionId);
      setAnalysisData(response.data);
    } catch (error) {
      toast.error('Analysis not ready yet. Please wait a moment.');
      setSelectedSubmission(null);
    } finally {
      setLoadingAnalysis(false);
    }
  };

  const handleSubmit = async () => {
    if (!code.trim()) {
      toast.error('Please write some code first');
      return;
    }

    setSubmitting(true);
    try {
      const response = await codeAPI.submit({
        language,
        source_code: code,
      });

      toast.success('Code submitted for analysis!');
      
      const newSubmission = {
        id: response.data.submission_id,
        language,
        timestamp: new Date().toISOString(),
        status: 'analyzing',
      };
      
      setSubmissions([newSubmission, ...submissions]);
      setCode('// Write your code here...\n\n');
    } catch (error) {
      toast.error(error.response?.data?.error || 'Failed to submit code');
    } finally {
      setSubmitting(false);
    }
  };

  const selectedLangColor = LANGUAGES.find(l => l.value === language)?.color || 'from-purple-500 to-cyan-500';

  return (
    <div className="min-h-screen bg-gray-950">
      <Navbar />
      
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
          {/* Code Editor - Main Column */}
          <div className="lg:col-span-2">
            <div className="glass rounded-2xl overflow-hidden border border-gray-800/50 shadow-2xl">
              <div className={`bg-gradient-to-r ${selectedLangColor} px-6 py-4`}>
                <div className="flex items-center justify-between">
                  <div className="flex items-center gap-3">
                    <div className="p-2 bg-white/20 rounded-lg backdrop-blur">
                      <Code2 className="w-5 h-5 text-white" />
                    </div>
                    <h2 className="text-xl font-semibold text-white">Code Editor</h2>
                  </div>
                  
                  <select
                    value={language}
                    onChange={(e) => setLanguage(e.target.value)}
                    className="bg-white/20 backdrop-blur text-white rounded-lg px-4 py-2 font-medium focus:outline-none focus:ring-2 focus:ring-white/50 cursor-pointer hover:bg-white/30 transition-all"
                  >
                    {LANGUAGES.map((lang) => (
                      <option key={lang.value} value={lang.value} className="text-gray-900 bg-gray-100">
                        {lang.label}
                      </option>
                    ))}
                  </select>
                </div>
              </div>

              <div className="h-[500px] bg-[#1e1e1e]">
                <Editor
                  height="100%"
                  language={language}
                  value={code}
                  onChange={(value) => setCode(value || '')}
                  theme="vs-dark"
                  options={{
                    minimap: { enabled: false },
                    fontSize: 14,
                    lineNumbers: 'on',
                    scrollBeyondLastLine: false,
                    automaticLayout: true,
                    fontFamily: 'Fira Code, Consolas, monospace',
                    fontLigatures: true,
                  }}
                />
              </div>

              <div className="px-6 py-4 bg-gray-900/50 border-t border-gray-800/50">
                <button
                  onClick={handleSubmit}
                  disabled={submitting}
                  className={`w-full bg-gradient-to-r ${selectedLangColor} text-white py-3 rounded-xl font-medium hover:scale-[1.02] transition-all disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center gap-2 glow-purple`}
                >
                  {submitting ? (
                    <>
                      <Loader2 className="w-5 h-5 animate-spin" />
                      Submitting...
                    </>
                  ) : (
                    <>
                      <Play className="w-5 h-5" />
                      Submit for Analysis
                    </>
                  )}
                </button>
              </div>
            </div>
          </div>

          {/* Sidebar - Submissions History */}
          <div className="lg:col-span-1">
            <div className="glass rounded-2xl p-6 border border-gray-800/50 shadow-2xl">
              <div className="flex items-center gap-2 mb-4">
                <Sparkles className="w-5 h-5 text-purple-400" />
                <h3 className="text-lg font-semibold text-white">Recent Submissions</h3>
              </div>
              
              {submissions.length === 0 ? (
                <div className="text-center py-12">
                  <div className="w-16 h-16 bg-gradient-to-br from-purple-500/20 to-cyan-500/20 rounded-2xl flex items-center justify-center mx-auto mb-4 border border-purple-500/30">
                    <Code2 className="w-8 h-8 text-purple-400" />
                  </div>
                  <p className="text-gray-400 font-medium">No submissions yet</p>
                  <p className="text-sm text-gray-500 mt-1">Submit your first code!</p>
                </div>
              ) : (
                <div className="space-y-3 max-h-[500px] overflow-y-auto pr-2">
                  {submissions.map((submission) => (
                    <div
                      key={submission.id}
                      onClick={() => viewAnalysis(submission.id)}
                      className="p-4 bg-gray-900/50 border border-gray-800/50 rounded-xl hover:border-purple-500/50 hover:bg-gray-800/50 transition-all cursor-pointer group"
                    >
                      <div className="flex items-center justify-between mb-2">
                        <span className="inline-flex items-center px-3 py-1 rounded-lg text-xs font-medium bg-purple-500/20 text-purple-300 border border-purple-500/30">
                          {submission.language}
                        </span>
                        {submission.status === 'analyzing' ? (
                          <Loader2 className="w-4 h-4 text-cyan-400 animate-spin" />
                        ) : submission.status === 'completed' ? (
                          <CheckCircle className="w-4 h-4 text-green-400" />
                        ) : (
                          <AlertCircle className="w-4 h-4 text-red-400" />
                        )}
                      </div>
                      <p className="text-xs text-gray-500 mb-2">
                        {new Date(submission.timestamp).toLocaleString()}
                      </p>
                      <p className="text-xs text-purple-400 font-medium group-hover:text-purple-300 transition-colors">
                        Click to view analysis →
                      </p>
                    </div>
                  ))}
                </div>
              )}

              <button
                onClick={() => navigate('/recommendations')}
                className="w-full mt-6 bg-gradient-to-r from-purple-600 to-pink-600 hover:from-purple-500 hover:to-pink-500 text-white py-3 rounded-xl font-medium transition-all text-sm flex items-center justify-center gap-2 glow-purple transform hover:scale-[1.02]"
              >
                <Sparkles className="w-4 h-4" />
                View Recommendations
              </button>
            </div>
          </div>
        </div>

        {/* Analysis Modal */}
        {selectedSubmission && (
          <div className="fixed inset-0 bg-black/80 backdrop-blur-sm flex items-center justify-center p-4 z-50 animate-in fade-in duration-200" onClick={() => setSelectedSubmission(null)}>
            <div className="glass rounded-2xl shadow-2xl max-w-2xl w-full max-h-[80vh] overflow-y-auto border border-gray-800/50 animate-in zoom-in duration-200" onClick={(e) => e.stopPropagation()}>
              <div className="bg-gradient-to-r from-purple-600 to-cyan-600 px-6 py-4 flex items-center justify-between">
                <div className="flex items-center gap-3">
                  <div className="p-2 bg-white/20 rounded-lg backdrop-blur">
                    <Zap className="w-5 h-5 text-white" />
                  </div>
                  <h3 className="text-xl font-semibold text-white">Analysis Results</h3>
                </div>
                <button onClick={() => setSelectedSubmission(null)} className="text-white hover:bg-white/20 rounded-lg p-2 transition-colors">
                  <X className="w-5 h-5" />
                </button>
              </div>

              <div className="p-6">
                {loadingAnalysis ? (
                  <div className="flex flex-col items-center justify-center py-12">
                    <Loader2 className="w-12 h-12 animate-spin text-purple-400 mb-4" />
                    <p className="text-gray-400">Analyzing your code...</p>
                  </div>
                ) : analysisData ? (
                  <div className="space-y-6">
                    {/* Complexity */}
                    <div className="grid grid-cols-2 gap-4">
                      <div className="bg-gradient-to-br from-blue-500/10 to-cyan-500/10 border border-blue-500/30 rounded-xl p-4">
                        <div className="flex items-center gap-2 mb-2">
                          <Clock className="w-5 h-5 text-cyan-400" />
                          <h4 className="font-semibold text-cyan-300">Time Complexity</h4>
                        </div>
                        <p className="text-2xl font-bold text-cyan-400">{analysisData.time_complexity || 'N/A'}</p>
                      </div>
                      <div className="bg-gradient-to-br from-purple-500/10 to-pink-500/10 border border-purple-500/30 rounded-xl p-4">
                        <div className="flex items-center gap-2 mb-2">
                          <Zap className="w-5 h-5 text-purple-400" />
                          <h4 className="font-semibold text-purple-300">Space Complexity</h4>
                        </div>
                        <p className="text-2xl font-bold text-purple-400">{analysisData.space_complexity || 'N/A'}</p>
                      </div>
                    </div>

                    {/* Patterns */}
                    {analysisData.patterns && analysisData.patterns.length > 0 && (
                      <div className="bg-gradient-to-br from-green-500/10 to-emerald-500/10 border border-green-500/30 rounded-xl p-4">
                        <h4 className="font-semibold text-green-300 mb-3 flex items-center gap-2">
                          <Sparkles className="w-5 h-5" />
                          Detected Patterns
                        </h4>
                        <div className="flex flex-wrap gap-2">
                          {analysisData.patterns.map((pattern, idx) => (
                            <span key={idx} className="inline-flex items-center px-3 py-1.5 rounded-lg text-sm font-medium bg-green-500/20 text-green-300 border border-green-500/30">
                              {pattern}
                            </span>
                          ))}
                        </div>
                      </div>
                    )}

                    {/* Issues */}
                    {analysisData.issues && (
                      <div className="bg-gradient-to-br from-yellow-500/10 to-orange-500/10 border border-yellow-500/30 rounded-xl p-4">
                        <h4 className="font-semibold text-yellow-300 mb-2 flex items-center gap-2">
                          <AlertCircle className="w-5 h-5" />
                          Issues
                        </h4>
                        <p className="text-yellow-200">{analysisData.issues}</p>
                      </div>
                    )}

                    {/* Metadata */}
                    <div className="bg-gray-900/50 rounded-xl p-4 text-sm text-gray-400 border border-gray-800/50">
                      <p><strong className="text-gray-300">Analyzed at:</strong> {analysisData.created_at}</p>
                      <p><strong className="text-gray-300">Submission ID:</strong> {analysisData.submission_id}</p>
                    </div>
                  </div>
                ) : (
                  <div className="text-center py-12">
                    <AlertCircle className="w-12 h-12 text-gray-600 mx-auto mb-3" />
                    <p className="text-gray-400">Analysis not available</p>
                  </div>
                )}
              </div>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}
