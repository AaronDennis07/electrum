import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

const SessionListPage = () => {
  const [sessions, setSessions] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    fetch('http://localhost:8000/session')
      .then(response => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json();
      })
      .then(data => {
        // Add random start times for upcoming sessions
        const enhancedData = data.map(session => {
          if (session.status === 'upcoming') {
            const randomFutureDate = new Date(Date.now() + Math.random() * 10 * 24 * 60 * 60 * 1000);
            return { ...session, startTime: randomFutureDate };
          }
          return session;
        });
        setSessions(enhancedData);
        setLoading(false);
      })
      .catch(error => {
        setError('Failed to fetch sessions');
        setLoading(false);
      });
  }, []);

  const getStatusColor = (status) => {
    switch (status) {
      case 'open':
        return 'bg-green-500';
      case 'closed':
        return 'bg-red-500';
      case 'upcoming':
        return 'bg-yellow-500';
      default:
        return 'bg-gray-500';
    }
  };

  const handleApply = (sessionName) => {
    navigate(`/enroll/${sessionName}`);
  };

  const renderProgressBar = (applied, total) => (
    <div className="w-full bg-gray-200 rounded-full h-2.5 mb-4">
      <div 
        className="bg-blue-600 h-2.5 rounded-full" 
        style={{ width: `${(applied / total) * 100}%` }}
      ></div>
    </div>
  );

  const renderSessionCard = (session) => (
    <div key={session.ID} className="bg-white shadow-md rounded-lg p-6 mb-4">
      <h3 className="text-xl font-semibold mb-2">{session.name}</h3>
      <p className="text-gray-600 mb-2">Type: {session.session_type}</p>
      <div className={`inline-block px-2 py-1 rounded-full text-white text-sm ${getStatusColor(session.status)} mb-2`}>
        {session.status}
      </div>
      {session.status === 'open' && (
        <>
          {renderProgressBar(session.applied_students, session.total_students)}
          <p className="text-gray-600 mb-2">
            {session.applied_students} / {session.total_students} students
          </p>
          <button 
            onClick={() => handleApply(session.name)}
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded mt-2"
          >
            Apply Now
          </button>
        </>
      )}
      {session.status === 'upcoming' && (
        <p className="text-gray-600 mt-2">
          Starts in: {session.startTime.toLocaleString()}
        </p>
      )}
      <p className="text-gray-500 text-sm mt-4">
        Created: {new Date(session.CreatedAt).toLocaleDateString()}
      </p>
    </div>
  );

  const renderSessionSection = (title, sessionList) => (
    <div className="mb-8">
      <h2 className="text-2xl font-bold mb-4">{title}</h2>
      {sessionList.length > 0 ? (
        sessionList.map(renderSessionCard)
      ) : (
        <p className="text-gray-500">No {title.toLowerCase()} sessions available.</p>
      )}
    </div>
  );

  if (loading) return <div className="text-center mt-8">Loading...</div>;
  if (error) return <div className="text-center mt-8 text-red-500">{error}</div>;

  const liveSessions = sessions.filter(session => session.status === 'open');
  const upcomingSessions = sessions.filter(session => session.status === 'upcoming');
  const closedSessions = sessions.filter(session => session.status === 'closed');

  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-3xl font-bold mb-8">Session List</h1>
      {renderSessionSection('Live', liveSessions)}
      {renderSessionSection('Upcoming', upcomingSessions)}
      {renderSessionSection('Closed', closedSessions)}
    </div>
  );
};

export default SessionListPage;