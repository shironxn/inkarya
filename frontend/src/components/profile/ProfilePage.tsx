"use client"

import { useEffect, useState } from 'react';
import { useStackApp } from '@stackframe/stack';

interface User {
  id: string;
  name: string;
  email: string;
  age: number;
  location: string;
  phone: string;
  status: string;
  availability: string;
  bio: string;
  skills: string[];
}

export default function ProfilePage() {
  const app = useStackApp();
  const user = app.useUser();
  const [profile, setProfile] = useState<User | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchUserProfile = async () => {
      try {
        const auth = await user?.getAuthJson();
        if (!auth) {
          throw new Error('Not authenticated');
        }

        const response = await fetch('/api/profile', {
          headers: {
            'Authorization': `Bearer ${auth.accessToken}`,
            'Content-Type': 'application/json'
          }
        });
        
        if (!response.ok) {
          throw new Error('Failed to fetch profile');
        }
        
        const data = await response.json();
        setProfile(data);
        setLoading(false);
      } catch {
        setError('Failed to load user profile');
        setLoading(false);
      }
    };

    if (user) {
      fetchUserProfile();
    }
  }, [user]);

  if (loading) {
    return <div className="flex justify-center items-center min-h-screen">Loading...</div>;
  }

  if (error) {
    return <div className="text-red-500 text-center">{error}</div>;
  }

  if (!profile) {
    return <div className="text-center">No user data found</div>;
  }

  return (
    <div className="max-w-4xl mx-auto px-4 py-8">
      <div className="bg-white rounded-lg shadow-lg p-6">
        <div className="flex items-center space-x-4 mb-6">
          <div className="w-24 h-24 bg-gray-200 rounded-full"></div>
          <div>
            <h1 className="text-2xl font-bold">{profile.name}</h1>
            <p className="text-gray-600">{profile.email}</p>
          </div>
        </div>

        <div className="grid grid-cols-2 gap-6">
          <div>
            <h2 className="text-lg font-semibold mb-2">Personal Information</h2>
            <div className="space-y-2">
              <div>
                <span className="text-gray-600">Age:</span> {profile.age}
              </div>
              <div>
                <span className="text-gray-600">Location:</span> {profile.location}
              </div>
              <div>
                <span className="text-gray-600">Phone:</span> {profile.phone}
              </div>
            </div>
          </div>

          <div>
            <h2 className="text-lg font-semibold mb-2">Work Details</h2>
            <div className="space-y-2">
              <div>
                <span className="text-gray-600">Status:</span> {profile.status}
              </div>
              <div>
                <span className="text-gray-600">Availability:</span> {profile.availability}
              </div>
            </div>
          </div>
        </div>

        <div className="mt-6">
          <h2 className="text-lg font-semibold mb-2">Bio</h2>
          <p className="text-gray-700">{profile.bio}</p>
        </div>

        <div className="mt-6">
          <h2 className="text-lg font-semibold mb-2">Skills</h2>
          <div className="flex flex-wrap gap-2">
            {profile.skills.map((skill, index) => (
              <span
                key={index}
                className="px-3 py-1 bg-blue-100 text-blue-800 rounded-full text-sm"
              >
                {skill}
              </span>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
} 