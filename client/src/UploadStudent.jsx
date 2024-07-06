import React, { useState } from 'react';
import toast, { Toaster } from 'react-hot-toast';

const UploadStudent = () => {

  const [file, setFile] = useState(null);
  const [isSubmitting, setIsSubmitting] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setIsSubmitting(true);

    const formData = new FormData();
    formData.append('student', file);


    try {
      const response = await fetch('http://localhost:8000/upload/student', {
        method: 'POST',
        body: formData,
      });

      if (!response.ok) {
        throw new Error('Failed to upload student');
      }

      toast.success('Students added successfully!');
      setFile(null);
    } catch (error) {
      toast.error(`Error adding students: ${error.message}`);
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <div className="container mx-auto p-4 bg-gray-100">
      <Toaster position="top-right" />
      <h1 className="text-3xl font-bold mb-6 text-center text-indigo-800">Add students</h1>
      <form onSubmit={handleSubmit} className="max-w-md mx-auto bg-white p-6 rounded-lg shadow-md">
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
          {isSubmitting ? 'Adding students...' : 'Upload students'}
        </button>
      </form>
    </div>
  );
};

export default UploadStudent;