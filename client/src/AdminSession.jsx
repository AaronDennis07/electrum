import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";

const AdminSessionPage = () => {
  const [sessions, setSessions] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    fetchSessions();
  }, []);

  const fetchSessions = () => {
    fetch("http://localhost:8000/session")
      .then((response) => {
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        return response.json();
      })
      .then((data) => {
        // Sort sessions by creation date in descending order
        const sortedSessions = data.sort(
          (a, b) => new Date(b.CreatedAt) - new Date(a.CreatedAt),
        );
        setSessions(sortedSessions);
        setLoading(false);
      })
      .catch((error) => {
        setError("Failed to fetch sessions");
        setLoading(false);
      });
  };

  const handleStartStop = (sessionName, action) => {
    fetch(`http://localhost:8000/session/${sessionName}/${action}`, {
      method: "POST",
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Failed to update session status");
        }
        return response.json();
      })
      .then(() => {
        // Refresh the session list after successful update
        fetchSessions();
      })
      .catch((error) => {
        setError(`Failed to ${action} session: ${error.message}`);
      });
  };

  const handleDetails = (sessionName) => {
    navigate(`/admin/session/${sessionName}`);
  };
  const getStatusColor = (status) => {
    switch (status) {
      case "open":
        return "bg-green-500";
      case "closed":
        return "bg-red-500";
      case "upcoming":
        return "bg-yellow-500";
      default:
        return "bg-gray-500";
    }
  };

  const renderSessionCard = (session) => (
    <div key={session.ID} className="bg-white shadow-md rounded-lg p-6 mb-4">
      <h3 className="text-xl font-semibold mb-2" onClick={() => handleDetails}>{session.name}</h3>
      <p className="text-gray-600 mb-2">Type: {session.session_type}</p>
      <p className="text-gray-600 mb-2">
        Total Students: {session.total_students}
      </p>
      <p className="text-gray-600 mb-2">
        Applied Students: {session.applied_students}
      </p>
      <div
        className={`inline-block px-2 py-1 rounded-full text-white text-sm ${getStatusColor(session.status)} mb-2`}
      >
        {session.status}
      </div>
      <p className="text-gray-500 text-sm mb-4">
        Created: {new Date(session.CreatedAt).toLocaleString()}
      </p>
      <div className="flex space-x-2">
        <button
          onClick={() => handleStartStop(session.name, "start")}
          className="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded"
          disabled={session.status === "open"}
        >
          Start
        </button>
        <button
          onClick={() => handleStartStop(session.name, "stop")}
          className="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded"
          disabled={session.status === "closed"}
        >
          Stop
        </button>
        <button
          onClick={() => handleDetails(session.name)}
          className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
          disabled={session.status === "closed"}
        >
          Details
        </button>
      </div>
    </div>
  );

  if (loading) return <div className="text-center mt-8">Loading...</div>;
  if (error)
    return <div className="text-center mt-8 text-red-500">{error}</div>;

  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-3xl font-bold mb-8">Admin Session Management</h1>
      <button
        className="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded"
        onClick={() => {
          navigate("/admin/create");
        }}
      >
        Create
      </button>
      <button
        className="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded mx-2"
        onClick={() => {
          navigate("/admin/upload");
        }}
      >
        Upload
      </button>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {sessions.map(renderSessionCard)}
      </div>
    </div>
  );
};

export default AdminSessionPage;
