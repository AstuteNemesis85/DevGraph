import { Code2, Info, Star, TrendingUp, Users, X, Sparkles, Zap } from 'lucide-react';
import { useEffect, useState } from 'react';
import toast from 'react-hot-toast';
import Navbar from '../components/Navbar';
import { useAuth } from '../context/AuthContext';
import { graphAPI } from '../services/api';

export default function Recommendations() {
  const { user } = useAuth();
  const [recommendations, setRecommendations] = useState([]);
  const [loading, setLoading] = useState(true);
  const [building, setBuilding] = useState(false);
  const [error, setError] = useState(null);
  const [selectedRec, setSelectedRec] = useState(null);

  useEffect(() => {
    fetchRecommendations();
  }, []);

  const fetchRecommendations = async () => {
    try {
      const response = await graphAPI.getRecommendations();
      setRecommendations(response.data || []);
      setError(null);
    } catch (error) {
      console.error('Recommendations error:', error);
      setRecommendations([]);
      setError(error.message);
    } finally {
      setLoading(false);
    }
  };

  const buildGraph = async () => {
    setBuilding(true);
    try {
      await graphAPI.buildGraph();
      toast.success('Graph built! Refreshing recommendations...');
      setLoading(true);
      await fetchRecommendations();
    } catch (error) {
      console.error('Build graph error:', error);
      toast.error('Failed to build graph');
    } finally {
      setBuilding(false);
    }
  };

  const getSimilarityColor = (similarity) => {
    if (similarity >= 0.8) return 'from-green-500/20 to-emerald-500/20 border-green-500/30 text-green-300';
    if (similarity >= 0.6) return 'from-blue-500/20 to-cyan-500/20 border-blue-500/30 text-blue-300';
    if (similarity >= 0.4) return 'from-yellow-500/20 to-orange-500/20 border-yellow-500/30 text-yellow-300';
    return 'from-gray-500/20 to-gray-600/20 border-gray-500/30 text-gray-300';
  };

  const getSimilarityLabel = (similarity) => {
    if (similarity >= 0.8) return 'Very High';
    if (similarity >= 0.6) return 'High';
    if (similarity >= 0.4) return 'Medium';
    return 'Low';
  };

  const getSimilarityGlow = (similarity) => {
    if (similarity >= 0.8) return 'hover:shadow-lg hover:shadow-green-500/20';
    if (similarity >= 0.6) return 'hover:shadow-lg hover:shadow-blue-500/20';
    if (similarity >= 0.4) return 'hover:shadow-lg hover:shadow-yellow-500/20';
    return 'hover:shadow-lg hover:shadow-gray-500/20';
  };

  return (
    <div className="min-h-screen bg-gray-950">
      <Navbar />
      
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div className="mb-8">
          <div className="flex items-center gap-3 mb-2">
            <div className="p-3 bg-gradient-to-br from-purple-500 to-pink-500 rounded-xl glow-purple">
              <Users className="w-6 h-6 text-white" />
            </div>
            <h1 className="text-3xl font-bold bg-gradient-to-r from-purple-400 via-pink-400 to-purple-400 bg-clip-text text-transparent">
              Developer Recommendations
            </h1>
          </div>
          <p className="text-gray-400 ml-16">
            Discover developers with similar coding patterns and interests
          </p>
        </div>

        {loading ? (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {[1, 2, 3].map((i) => (
              <div key={i} className="glass rounded-2xl p-6 border border-gray-800/50 animate-pulse">
                <div className="h-20 w-20 bg-gray-800 rounded-full mx-auto mb-4"></div>
                <div className="h-4 bg-gray-800 rounded w-3/4 mx-auto mb-2"></div>
                <div className="h-3 bg-gray-800 rounded w-1/2 mx-auto"></div>
              </div>
            ))}
          </div>
        ) : recommendations.length === 0 ? (
          <div className="space-y-6">
            <div className="glass rounded-2xl p-12 text-center border border-gray-800/50">
              <div className="w-24 h-24 bg-gradient-to-br from-purple-500/20 to-pink-500/20 rounded-full flex items-center justify-center mx-auto mb-6 border border-purple-500/30">
                <Users className="w-12 h-12 text-purple-400" />
              </div>
              <h3 className="text-2xl font-bold text-white mb-3">No Recommendations Yet</h3>
              <p className="text-gray-400 mb-2 max-w-2xl mx-auto">
                Recommendations will appear when other developers with similar coding patterns join the platform.
              </p>
              <p className="text-sm text-gray-500 mb-8">
                If you have multiple accounts with code submissions, click below to analyze similarities!
              </p>
              <div className="flex items-center justify-center gap-4">
                <button
                  onClick={buildGraph}
                  disabled={building}
                  className="inline-flex items-center gap-2 bg-gradient-to-r from-green-600 to-emerald-600 hover:from-green-500 hover:to-emerald-500 text-white px-6 py-3 rounded-xl font-medium transition-all disabled:opacity-50 disabled:cursor-not-allowed transform hover:scale-105 shadow-lg"
                >
                  {building ? (
                    <>
                      <TrendingUp className="w-5 h-5 animate-pulse" />
                      Building Graph...
                    </>
                  ) : (
                    <>
                      <TrendingUp className="w-5 h-5" />
                      Build Similarity Graph
                    </>
                  )}
                </button>
                <a
                  href="/dashboard"
                  className="inline-flex items-center gap-2 bg-gradient-to-r from-purple-600 to-cyan-600 hover:from-purple-500 hover:to-cyan-500 text-white px-6 py-3 rounded-xl font-medium transition-all transform hover:scale-105 shadow-lg"
                >
                  <Code2 className="w-5 h-5" />
                  Submit More Code
                </a>
              </div>
            </div>

            {/* How it Works Section */}
            <div className="glass rounded-2xl p-6 border border-blue-500/30">
              <div className="flex items-start gap-4">
                <div className="p-3 bg-gradient-to-br from-blue-500 to-cyan-500 rounded-xl glow">
                  <TrendingUp className="w-6 h-6 text-white" />
                </div>
                <div className="flex-1">
                  <h3 className="font-bold text-white text-xl mb-4">How Recommendations Work</h3>
                  <div className="space-y-4 text-sm text-gray-300">
                    <div className="flex items-start gap-3">
                      <div className="w-8 h-8 rounded-full bg-gradient-to-br from-blue-600 to-cyan-600 text-white flex items-center justify-center flex-shrink-0 text-sm font-bold shadow-lg">1</div>
                      <p><strong className="text-white">Submit Code:</strong> Your code is analyzed for algorithm patterns, complexity, and style</p>
                    </div>
                    <div className="flex items-start gap-3">
                      <div className="w-8 h-8 rounded-full bg-gradient-to-br from-purple-600 to-pink-600 text-white flex items-center justify-center flex-shrink-0 text-sm font-bold shadow-lg">2</div>
                      <p><strong className="text-white">Pattern Detection:</strong> System identifies sorting, searching, dynamic programming, and other patterns</p>
                    </div>
                    <div className="flex items-start gap-3">
                      <div className="w-8 h-8 rounded-full bg-gradient-to-br from-pink-600 to-red-600 text-white flex items-center justify-center flex-shrink-0 text-sm font-bold shadow-lg">3</div>
                      <p><strong className="text-white">Similarity Matching:</strong> When other developers submit code, similarities are calculated</p>
                    </div>
                    <div className="flex items-start gap-3">
                      <div className="w-8 h-8 rounded-full bg-gradient-to-br from-green-600 to-emerald-600 text-white flex items-center justify-center flex-shrink-0 text-sm font-bold shadow-lg">4</div>
                      <p><strong className="text-white">Get Recommendations:</strong> Discover developers with matching coding styles and interests</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            {/* Stats Card */}
            <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
              <div className="glass rounded-2xl p-6 text-center border border-blue-500/30 group hover:border-blue-400/50 transition-all">
                <div className="w-12 h-12 bg-gradient-to-br from-blue-500 to-cyan-500 rounded-xl mx-auto mb-3 flex items-center justify-center group-hover:scale-110 transition-transform">
                  <Code2 className="w-6 h-6 text-white" />
                </div>
                <p className="text-2xl font-bold text-white mb-1">Submit Code</p>
                <p className="text-sm text-gray-400">Build your profile</p>
              </div>
              <div className="glass rounded-2xl p-6 text-center border border-purple-500/30 group hover:border-purple-400/50 transition-all">
                <div className="w-12 h-12 bg-gradient-to-br from-purple-500 to-pink-500 rounded-xl mx-auto mb-3 flex items-center justify-center group-hover:scale-110 transition-transform">
                  <TrendingUp className="w-6 h-6 text-white" />
                </div>
                <p className="text-2xl font-bold text-white mb-1">Get Analyzed</p>
                <p className="text-sm text-gray-400">Patterns detected</p>
              </div>
              <div className="glass rounded-2xl p-6 text-center border border-pink-500/30 group hover:border-pink-400/50 transition-all">
                <div className="w-12 h-12 bg-gradient-to-br from-pink-500 to-red-500 rounded-xl mx-auto mb-3 flex items-center justify-center group-hover:scale-110 transition-transform">
                  <Users className="w-6 h-6 text-white" />
                </div>
                <p className="text-2xl font-bold text-white mb-1">Find Matches</p>
                <p className="text-sm text-gray-400">Connect with devs</p>
              </div>
            </div>
          </div>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {recommendations.map((rec) => {
              const otherUserId = rec.user_a === user?.id ? rec.user_b : rec.user_a;
              const sharedPatterns = rec.shared_patterns || [];
              
              return (
                <div
                  key={rec.id}
                  className={`glass rounded-2xl p-6 border border-gray-800/50 cursor-pointer transition-all transform hover:scale-105 ${getSimilarityGlow(rec.similarity)}`}
                  onClick={() => setSelectedRec(rec)}
                >
                  <div className="text-center">
                    <div className="w-24 h-24 bg-gradient-to-br from-purple-500 to-cyan-600 rounded-full flex items-center justify-center mx-auto mb-4 shadow-lg">
                      <span className="text-3xl font-bold text-white">
                        {String(otherUserId || '').substring(0, 2).toUpperCase()}
                      </span>
                    </div>
                    
                    <h3 className="text-lg font-semibold text-white mb-2">
                      Developer #{String(otherUserId || '').substring(0, 8)}
                    </h3>
                    
                    <div className="flex items-center justify-center gap-2 mb-4">
                      <Sparkles className="w-4 h-4 text-purple-400" />
                      <span className="text-sm text-gray-400">
                        {rec.total_patterns || 0} shared patterns
                      </span>
                    </div>

                    {sharedPatterns.length > 0 && (
                      <div className="flex flex-wrap gap-2 justify-center mb-4">
                        {sharedPatterns.slice(0, 3).map((pattern, idx) => (
                          <span key={idx} className="px-3 py-1 bg-purple-500/20 text-purple-300 text-xs rounded-lg border border-purple-500/30 font-medium">
                            {pattern}
                          </span>
                        ))}
                        {sharedPatterns.length > 3 && (
                          <span className="px-3 py-1 bg-gray-700/50 text-gray-300 text-xs rounded-lg border border-gray-600/30 font-medium">
                            +{sharedPatterns.length - 3}
                          </span>
                        )}
                      </div>
                    )}

                    <div className={`inline-flex items-center gap-2 px-4 py-2 rounded-xl bg-gradient-to-r ${getSimilarityColor(rec.similarity)} border`}>
                      <Star className="w-4 h-4" />
                      <span className="font-bold text-sm">
                        {getSimilarityLabel(rec.similarity)} Match
                      </span>
                      <span className="text-xs opacity-75">
                        ({(rec.similarity * 100).toFixed(0)}%)
                      </span>
                    </div>

                    <button 
                      className="mt-4 text-purple-400 hover:text-purple-300 text-sm font-medium flex items-center gap-1 mx-auto transition-colors"
                      onClick={(e) => {
                        e.stopPropagation();
                        setSelectedRec(rec);
                      }}
                    >
                      <Info className="w-4 h-4" />
                      View Details
                    </button>

                    {rec.created_at && (
                      <p className="text-xs text-gray-500 mt-3">
                        Updated {new Date(rec.created_at).toLocaleDateString()}
                      </p>
                    )}
                  </div>
                </div>
              );
            })}
          </div>
        )}

        {/* Rebuild Graph Button */}
        <div className="mt-8 glass rounded-2xl p-6 border border-green-500/30">
          <div className="flex items-center justify-between">
            <div className="flex items-start gap-4">
              <div className="p-3 bg-gradient-to-br from-green-500 to-emerald-500 rounded-xl shadow-lg">
                <TrendingUp className="w-6 h-6 text-white" />
              </div>
              <div>
                <h3 className="font-bold text-white text-lg mb-2">Rebuild Similarity Graph</h3>
                <p className="text-sm text-gray-400">
                  Click here to recalculate similarities between all users. This will analyze all code submissions and update recommendations.
                </p>
              </div>
            </div>
            <button
              onClick={buildGraph}
              disabled={building}
              className="flex-shrink-0 inline-flex items-center gap-2 bg-gradient-to-r from-green-600 to-emerald-600 hover:from-green-500 hover:to-emerald-500 text-white px-6 py-3 rounded-xl font-medium transition-all disabled:opacity-50 disabled:cursor-not-allowed shadow-lg transform hover:scale-105"
            >
              {building ? (
                <>
                  <TrendingUp className="w-5 h-5 animate-pulse" />
                  Building...
                </>
              ) : (
                <>
                  <TrendingUp className="w-5 h-5" />
                  Build Graph
                </>
              )}
            </button>
          </div>
        </div>
      </div>

      {/* Detail Modal */}
      {selectedRec && (
        <div className="fixed inset-0 bg-black/80 backdrop-blur-sm flex items-center justify-center p-4 z-50" onClick={() => setSelectedRec(null)}>
          <div className="glass rounded-2xl shadow-2xl max-w-2xl w-full max-h-[90vh] overflow-y-auto border border-gray-800/50" onClick={(e) => e.stopPropagation()}>
            <div className="sticky top-0 bg-gradient-to-r from-purple-600 to-cyan-600 text-white p-6 flex items-center justify-between">
              <div className="flex items-center gap-3">
                <div className="w-12 h-12 bg-white/20 rounded-xl flex items-center justify-center backdrop-blur">
                  <Users className="w-6 h-6" />
                </div>
                <div>
                  <h2 className="text-xl font-bold">Developer Match Details</h2>
                  <p className="text-purple-100 text-sm">Similarity Analysis</p>
                </div>
              </div>
              <button onClick={() => setSelectedRec(null)} className="text-white hover:bg-white/20 p-2 rounded-xl transition-colors">
                <X className="w-5 h-5" />
              </button>
            </div>

            <div className="p-6 space-y-6">
              {/* Similarity Score */}
              <div className="text-center bg-gradient-to-br from-purple-500/10 to-cyan-500/10 rounded-2xl p-6 border border-purple-500/30">
                <div className={`inline-flex items-center gap-3 px-6 py-3 rounded-xl bg-gradient-to-r ${getSimilarityColor(selectedRec.similarity)} border shadow-lg`}>
                  <Star className="w-6 h-6" />
                  <span className="font-bold text-2xl">
                    {(selectedRec.similarity * 100).toFixed(1)}% Match
                  </span>
                </div>
                <p className="text-sm text-gray-400 mt-3">
                  {getSimilarityLabel(selectedRec.similarity)} similarity based on coding patterns
                </p>
              </div>

              {/* Shared Patterns */}
              {selectedRec.shared_patterns && selectedRec.shared_patterns.length > 0 && (
                <div className="bg-green-500/10 rounded-2xl p-5 border border-green-500/30">
                  <div className="flex items-center gap-2 mb-3">
                    <Zap className="w-5 h-5 text-green-400" />
                    <h3 className="font-bold text-green-300">Shared Patterns ({selectedRec.total_patterns})</h3>
                  </div>
                  <div className="flex flex-wrap gap-2">
                    {selectedRec.shared_patterns.map((pattern, idx) => (
                      <span key={idx} className="px-4 py-2 bg-green-500/20 text-green-300 rounded-xl font-medium text-sm border border-green-500/30">
                        {pattern}
                      </span>
                    ))}
                  </div>
                  <p className="text-sm text-gray-400 mt-3">
                    You both use these algorithm patterns in your code
                  </p>
                </div>
              )}

              {/* Your Patterns & Their Patterns */}
              {selectedRec.user_a_patterns && selectedRec.user_b_patterns && (
                <>
                  <div className="bg-blue-500/10 rounded-2xl p-5 border border-blue-500/30">
                    <div className="flex items-center gap-2 mb-3">
                      <Star className="w-5 h-5 text-blue-400" />
                      <h3 className="font-bold text-blue-300">Your Patterns</h3>
                    </div>
                    <div className="flex flex-wrap gap-2">
                      {(selectedRec.user_a === user?.id ? selectedRec.user_a_patterns : selectedRec.user_b_patterns).map((pattern, idx) => (
                        <span key={idx} className="px-3 py-1.5 bg-blue-500/20 text-blue-300 rounded-lg text-sm border border-blue-500/30">
                          {pattern}
                        </span>
                      ))}
                    </div>
                  </div>

                  <div className="bg-purple-500/10 rounded-2xl p-5 border border-purple-500/30">
                    <div className="flex items-center gap-2 mb-3">
                      <Users className="w-5 h-5 text-purple-400" />
                      <h3 className="font-bold text-purple-300">Their Patterns</h3>
                    </div>
                    <div className="flex flex-wrap gap-2">
                      {(selectedRec.user_a === user?.id ? selectedRec.user_b_patterns : selectedRec.user_a_patterns).map((pattern, idx) => (
                        <span key={idx} className="px-3 py-1.5 bg-purple-500/20 text-purple-300 rounded-lg text-sm border border-purple-500/30">
                          {pattern}
                        </span>
                      ))}
                    </div>
                  </div>
                </>
              )}

              {/* Info Box */}
              <div className="bg-cyan-500/10 border border-cyan-500/30 rounded-2xl p-4">
                <div className="flex items-start gap-3">
                  <Info className="w-5 h-5 text-cyan-400 flex-shrink-0 mt-0.5" />
                  <div className="text-sm text-gray-300">
                    <p className="font-semibold text-cyan-300 mb-2">How similarity is calculated:</p>
                    <p>
                      We use Weighted Jaccard similarity to compare algorithm patterns detected in your code submissions. 
                      A higher percentage means more common patterns between developers.
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
