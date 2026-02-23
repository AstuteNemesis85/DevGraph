import { Code2, LogOut, TrendingUp, User } from 'lucide-react';
import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';

export default function Navbar() {
  const { user, logout } = useAuth();
  const navigate = useNavigate();

  const handleLogout = () => {
    logout();
    navigate('/login');
  };

  return (
    <nav className="glass border-b border-gray-800 sticky top-0 z-50">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between items-center h-16">
          {/* Logo */}
          <Link to="/dashboard" className="flex items-center gap-3 group">
            <div className="p-2 bg-gradient-to-br from-purple-500 to-cyan-500 rounded-xl group-hover:scale-110 transition-transform duration-300 glow-purple">
              <Code2 className="w-6 h-6 text-white" />
            </div>
            <span className="text-2xl font-bold bg-gradient-to-r from-purple-400 via-cyan-400 to-purple-400 bg-clip-text text-transparent">
              DevGraph
            </span>
          </Link>

          {/* Navigation Links */}
          <div className="flex items-center gap-4">
            <Link
              to="/dashboard"
              className="flex items-center gap-2 px-4 py-2 rounded-xl bg-gray-800/50 hover:bg-gray-700/50 text-gray-300 hover:text-white transition-all duration-200 border border-gray-700/50 hover:border-purple-500/50"
            >
              <Code2 className="w-4 h-4" />
              <span className="font-medium">Dashboard</span>
            </Link>

            <Link
              to="/recommendations"
              className="flex items-center gap-2 px-4 py-2 rounded-xl bg-gradient-to-r from-purple-600/20 to-cyan-600/20 hover:from-purple-600/30 hover:to-cyan-600/30 text-purple-300 hover:text-purple-200 transition-all duration-200 border border-purple-500/30 hover:border-purple-400/50"
            >
              <TrendingUp className="w-4 h-4" />
              <span className="font-medium">Recommendations</span>
            </Link>

            {/* User Menu */}
            <div className="flex items-center gap-3 pl-3 border-l border-gray-700">
              {/* Clicking the user chip goes to /profile */}
              <Link
                to="/profile"
                className="flex items-center gap-2 px-3 py-2 bg-gray-800/50 hover:bg-gray-700/50 rounded-lg border border-gray-700/50 hover:border-purple-500/50 transition-all duration-200 group"
              >
                {/* Avatar: real image if set, gradient fallback otherwise */}
                <div className="w-8 h-8 rounded-lg overflow-hidden flex-shrink-0 ring-1 ring-purple-500/30 group-hover:ring-purple-400/60 transition-all">
                  {user?.avatar_url ? (
                    <img
                      src={user.avatar_url}
                      alt={user.username}
                      className="w-full h-full object-cover"
                      onError={(e) => { e.currentTarget.style.display = 'none'; }}
                    />
                  ) : (
                    <div className="w-full h-full bg-gradient-to-br from-purple-500 to-cyan-500 flex items-center justify-center">
                      <User className="w-4 h-4 text-white" />
                    </div>
                  )}
                </div>
                <span className="text-sm text-gray-300 font-medium group-hover:text-white transition-colors">
                  {user?.username || 'User'}
                </span>
              </Link>

              <button
                onClick={handleLogout}
                className="p-2 rounded-xl bg-red-500/10 hover:bg-red-500/20 text-red-400 hover:text-red-300 transition-all duration-200 border border-red-500/30 hover:border-red-400/50"
                title="Logout"
              >
                <LogOut className="w-5 h-5" />
              </button>
            </div>
          </div>
        </div>
      </div>
    </nav>
  );
}
