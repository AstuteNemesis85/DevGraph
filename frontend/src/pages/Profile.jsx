import { Camera, Check, Mail, Save, User, X } from 'lucide-react';
import { useEffect, useRef, useState } from 'react';
import toast from 'react-hot-toast';
import Navbar from '../components/Navbar';
import { useAuth } from '../context/AuthContext';
import { profileAPI } from '../services/api';

export default function Profile() {
  const { user, setUser } = useAuth();

  const [form, setForm] = useState({ username: '', bio: '', avatar_url: '' });
  const [preview, setPreview] = useState('');
  const [urlInput, setUrlInput] = useState('');
  const [showUrlInput, setShowUrlInput] = useState(false);
  const [saving, setSaving] = useState(false);
  const [imgError, setImgError] = useState(false);
  const urlRef = useRef(null);

  // Populate form from AuthContext user once loaded
  useEffect(() => {
    if (user) {
      setForm({
        username:   user.username   || '',
        bio:        user.bio        || '',
        avatar_url: user.avatar_url || '',
      });
      setPreview(user.avatar_url || '');
      setUrlInput(user.avatar_url || '');
    }
  }, [user]);

  // Focus the URL input when it appears
  useEffect(() => {
    if (showUrlInput && urlRef.current) urlRef.current.focus();
  }, [showUrlInput]);

  const handleChange = (e) => {
    setForm((prev) => ({ ...prev, [e.target.name]: e.target.value }));
  };

  // Apply the pasted URL as avatar preview
  const applyUrl = () => {
    setImgError(false);
    setForm((prev) => ({ ...prev, avatar_url: urlInput }));
    setPreview(urlInput);
    setShowUrlInput(false);
  };

  const cancelUrl = () => {
    setUrlInput(form.avatar_url);
    setShowUrlInput(false);
  };

  const handleSave = async () => {
    if (!form.username.trim()) {
      toast.error('Username cannot be empty');
      return;
    }
    setSaving(true);
    try {
      const res = await profileAPI.updateProfile({
        username:   form.username.trim(),
        bio:        form.bio,
        avatar_url: form.avatar_url,
      });
      // Update AuthContext so Navbar refreshes immediately
      setUser((prev) => ({ ...prev, ...res.data }));
      toast.success('Profile updated!');
    } catch (err) {
      toast.error(err.response?.data?.error || 'Failed to save profile');
    } finally {
      setSaving(false);
    }
  };

  const initials = (form.username || 'U').charAt(0).toUpperCase();

  return (
    <div className="min-h-screen bg-gray-950">
      <Navbar />

      <div className="max-w-2xl mx-auto px-4 py-12">

        {/* Page heading */}
        <div className="mb-8">
          <h1 className="text-3xl font-bold bg-gradient-to-r from-purple-400 to-cyan-400 bg-clip-text text-transparent">
            Profile
          </h1>
          <p className="text-gray-400 mt-1 text-sm">
            Manage your public identity on DevGraph
          </p>
        </div>

        <div className="glass rounded-2xl border border-gray-800 p-8 space-y-8">

          {/* ── Avatar section ─────────────────────────────────── */}
          <div className="flex items-start gap-6">
            {/* Avatar display */}
            <div className="relative flex-shrink-0">
              <div className="w-24 h-24 rounded-2xl overflow-hidden ring-2 ring-purple-500/40">
                {preview && !imgError ? (
                  <img
                    src={preview}
                    alt="avatar"
                    className="w-full h-full object-cover"
                    onError={() => setImgError(true)}
                  />
                ) : (
                  <div className="w-full h-full bg-gradient-to-br from-purple-500 to-cyan-500 flex items-center justify-center">
                    <span className="text-3xl font-bold text-white select-none">
                      {initials}
                    </span>
                  </div>
                )}
              </div>

              {/* Camera button */}
              <button
                onClick={() => { setShowUrlInput((v) => !v); }}
                className="absolute -bottom-2 -right-2 w-8 h-8 bg-purple-600 hover:bg-purple-500 rounded-lg flex items-center justify-center transition-colors shadow-lg"
                title="Change avatar"
              >
                <Camera className="w-4 h-4 text-white" />
              </button>
            </div>

            {/* Avatar info + URL input */}
            <div className="flex-1 min-w-0">
              <p className="text-white font-semibold text-lg truncate">{form.username || 'User'}</p>
              <p className="text-gray-400 text-sm mt-0.5">
                {user?.email || ''}
              </p>
              {user?.created_at && (
                <p className="text-gray-500 text-xs mt-1">Joined {user.created_at}</p>
              )}

              {/* URL input panel */}
              {showUrlInput && (
                <div className="mt-3 flex items-center gap-2">
                  <input
                    ref={urlRef}
                    type="url"
                    value={urlInput}
                    onChange={(e) => setUrlInput(e.target.value)}
                    onKeyDown={(e) => { if (e.key === 'Enter') applyUrl(); if (e.key === 'Escape') cancelUrl(); }}
                    placeholder="Paste image URL…"
                    className="flex-1 px-3 py-2 bg-gray-900 border border-gray-700 rounded-lg text-sm text-gray-200 placeholder-gray-600 focus:outline-none focus:border-purple-500 transition-colors"
                  />
                  <button
                    onClick={applyUrl}
                    className="p-2 bg-purple-600 hover:bg-purple-500 rounded-lg text-white transition-colors"
                    title="Apply"
                  >
                    <Check className="w-4 h-4" />
                  </button>
                  <button
                    onClick={cancelUrl}
                    className="p-2 bg-gray-700 hover:bg-gray-600 rounded-lg text-gray-300 transition-colors"
                    title="Cancel"
                  >
                    <X className="w-4 h-4" />
                  </button>
                </div>
              )}

              {!showUrlInput && (
                <button
                  onClick={() => setShowUrlInput(true)}
                  className="mt-2 text-xs text-purple-400 hover:text-purple-300 transition-colors"
                >
                  {form.avatar_url ? 'Change avatar URL' : '+ Set avatar URL'}
                </button>
              )}
            </div>
          </div>

          {/* Divider */}
          <div className="border-t border-gray-800" />

          {/* ── Username ──────────────────────────────────────── */}
          <div className="space-y-2">
            <label className="flex items-center gap-2 text-sm font-medium text-gray-300">
              <User className="w-4 h-4 text-purple-400" />
              Username
            </label>
            <input
              name="username"
              value={form.username}
              onChange={handleChange}
              maxLength={32}
              className="w-full px-4 py-3 bg-gray-900 border border-gray-700 rounded-xl text-gray-100 placeholder-gray-600 focus:outline-none focus:border-purple-500 focus:ring-1 focus:ring-purple-500/30 transition-all"
              placeholder="Your username"
            />
          </div>

          {/* ── Email (read-only) ─────────────────────────────── */}
          <div className="space-y-2">
            <label className="flex items-center gap-2 text-sm font-medium text-gray-300">
              <Mail className="w-4 h-4 text-cyan-400" />
              Email
              <span className="ml-1 text-xs text-gray-600 font-normal">(cannot change)</span>
            </label>
            <input
              value={user?.email || ''}
              readOnly
              className="w-full px-4 py-3 bg-gray-900/50 border border-gray-800 rounded-xl text-gray-500 cursor-not-allowed"
            />
          </div>

          {/* ── Bio ──────────────────────────────────────────── */}
          <div className="space-y-2">
            <label className="text-sm font-medium text-gray-300">Bio</label>
            <textarea
              name="bio"
              value={form.bio}
              onChange={handleChange}
              maxLength={200}
              rows={3}
              className="w-full px-4 py-3 bg-gray-900 border border-gray-700 rounded-xl text-gray-100 placeholder-gray-600 focus:outline-none focus:border-purple-500 focus:ring-1 focus:ring-purple-500/30 transition-all resize-none"
              placeholder="Tell other devs a bit about yourself…"
            />
            <p className="text-right text-xs text-gray-600">{form.bio.length}/200</p>
          </div>

          {/* ── Save button ──────────────────────────────────── */}
          <button
            onClick={handleSave}
            disabled={saving}
            className="w-full flex items-center justify-center gap-2 py-3 px-6 bg-gradient-to-r from-purple-600 to-cyan-600 hover:from-purple-500 hover:to-cyan-500 disabled:opacity-50 disabled:cursor-not-allowed text-white font-semibold rounded-xl transition-all duration-200 shadow-lg hover:shadow-purple-500/25"
          >
            {saving ? (
              <span className="w-5 h-5 border-2 border-white/40 border-t-white rounded-full animate-spin" />
            ) : (
              <Save className="w-5 h-5" />
            )}
            {saving ? 'Saving…' : 'Save Changes'}
          </button>

        </div>
      </div>
    </div>
  );
}
