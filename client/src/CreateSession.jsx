import React, { useState } from 'react';
import toast, { Toaster } from 'react-hot-toast';

const CreateSessionForm = () => {
  const [sessionName, setSessionName] = useState('');
  const [sessionType, setSessionType] = useState('open');
  const [file, setFile] = useState(null);
  const [isSubmitting, setIsSubmitting] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setIsSubmitting(true);

    const formData = new FormData();
    formData.append('file', file);

    const sessionData = {
      session: {
        name: sessionName,
        session_type: sessionType
      }
    };

    formData.append('data', JSON.stringify(sessionData));

    try {
      const response = await fetch('http://localhost:8000/session', {
        method: 'POST',
        body: formData,
      });

      if (!response.ok) {
        throw new Error('Failed to create session');
      }

      toast.success('Session created successfully!');
      setSessionName('');
      setSessionType('open');
      setFile(null);
    } catch (error) {
      toast.error(`Error creating session: ${error.message}`);
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <div className="container mx-auto p-4 bg-gray-100">
      <Toaster position="top-right" />
      <h1 className="text-3xl font-bold mb-6 text-center text-indigo-800">Create New Session</h1>
      <form onSubmit={handleSubmit} className="max-w-md mx-auto bg-white p-6 rounded-lg shadow-md">
        <div className="mb-4">
          <label htmlFor="sessionName" className="block text-gray-700 text-sm font-bold mb-2">
            Session Name:
          </label>
          <input
            type="text"
            id="sessionName"
            value={sessionName}
            onChange={(e) => setSessionName(e.target.value)}
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            required
          />
        </div>
        <div className="mb-4">
          <label htmlFor="sessionType" className="block text-gray-700 text-sm font-bold mb-2">
            Session Type:
          </label>
          <select
            id="sessionType"
            value={sessionType}
            onChange={(e) => setSessionType(e.target.value)}
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          >
            <option value="open">Open</option>
            <option value="professional">Professional</option>
          </select>
        </div>
        <div className="mb-6">
          <label htmlFor="file" className="block text-gray-700 text-sm font-bold mb-2">
            Upload Excel File:
          </label>
          <input
            type="file"
            id="file"
            onChange={(e) => setFile(e.target.files[0])}
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            accept=".xlsx, .xls"
            required
          />
        </div>
        <button
          type="submit"
          className="w-full bg-indigo-500 hover:bg-indigo-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline transition duration-300 ease-in-out"
          disabled={isSubmitting}
        >
          {isSubmitting ? 'Creating Session...' : 'Create Session'}
        </button>
      </form>
    </div>
  );
};

export default CreateSessionForm;