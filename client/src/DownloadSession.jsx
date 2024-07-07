import { useState, useEffect } from "react";
import toast, { Toaster } from "react-hot-toast";

const DownloadSession = () => {
  const [sessions, setSessions] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetchSessions().then(() => {
      console.log("success");
    });
    console.log(sessions);
  }, []);

  const fetchSessions = async () => {
    setLoading(true);
    try {
      const response = await fetch("http://localhost:8000/session");
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }
      const data = await response.json();
      setSessions(data);
      setLoading(false);
    } catch (error) {
      setError("Failed to fetch sessions");
      setLoading(false);
      toast.error("Failed to load sessions. Please try again later.");
    }
  };

  const handleDownload = async (sessionName) => {
    try {
      const response = await fetch(
        `http://localhost:8000/session/${sessionName}/excel`,
        {
          method: "GET",
          headers: {
            "Content-Type":
              "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
          },
        },
      );

      if (!response.ok) {
        throw new Error("Enrollment failed");
      }
      const blob = await response.blob();
      const url = window.URL.createObjectURL(blob);

      // Create a link element and trigger the download
      const link = document.createElement("a");
      link.href = url;
      link.setAttribute("download", `${sessionName}.xlsx`);
      document.body.appendChild(link);
      link.click();
      document.body.removeChild(link);
    } catch (error) {
      toast.error(`${error.message}`);
    }
  };

  if (loading) return <div className="text-center mt-8">Loading...</div>;
  if (error)
    return <div className="text-center mt-8 text-red-500">{error}</div>;

  return (
    <div className="container mx-auto p-4 bg-gray-100">
      <Toaster position="top-right" />
      <h1 className="text-3xl font-bold mb-6 text-center text-indigo-800">
        Session Details
      </h1>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {sessions.map((session) => (
          <div
            key={session.Id}
            className="bg-white rounded-lg shadow-lg overflow-hidden hover:shadow-xl transition-shadow duration-300 border border-indigo-100"
          >
            <div className="p-6">
              <h2 className="text-xl font-semibold mb-2 text-indigo-700">
                {session.name}
              </h2>
              <div className="mb-4">
                <div className="flex justify-between items-center mb-1">
                  <span className="text-sm font-medium text-indigo-600">
                    {session.session_type}
                  </span>
                </div>
              </div>
              <button
                className="w-full bg-indigo-500 hover:bg-indigo-700 text-white font-bold py-2 px-4 rounded transition duration-300 disabled:bg-gray-300 disabled:cursor-not-allowed"
                onClick={() => handleDownload(session.name)}
              >
                Download
              </button>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default DownloadSession;
