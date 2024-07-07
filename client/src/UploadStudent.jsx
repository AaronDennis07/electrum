import React, { useState, useRef } from "react";
import { toast, Toaster } from "react-hot-toast";
import { OutTable, ExcelRenderer } from 'react-excel-renderer';
import { Upload, File, X } from "lucide-react";

const UploadStudent = () => {
  const [file, setFile] = useState(null);
  const [fileName, setFileName] = useState("");
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [previewData, setPreviewData] = useState(null);
  const fileInputRef = useRef(null);

  const handleFileChange = (e) => {
    const selectedFile = e.target.files[0];
    if (selectedFile) {
      setFile(selectedFile);
      setFileName(selectedFile.name);

      ExcelRenderer(selectedFile, (err, resp) => {
        if (err) {
          console.error(err);
          toast.error("Error reading Excel file");
        } else {
          setPreviewData({
            cols: resp.cols.slice(0, 5),  // Limit to first 5 columns
            rows: resp.rows.slice(0, 21)  // Show header + first 20 rows
          });
        }
      });
    }
  };

  const handleRemoveFile = () => {
    setFile(null);
    setFileName("");
    setPreviewData(null);
    if (fileInputRef.current) {
      fileInputRef.current.value = "";
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setIsSubmitting(true);

    const formData = new FormData();
    formData.append("student", file);

    try {
      const response = await fetch("http://localhost:8000/api/v1/student/upload", {
        method: "POST",
        body: formData,
      });

      if (!response.ok) {
        throw new Error("Failed to upload student");
      }

      toast.success("Students added successfully!");
      handleRemoveFile();
    } catch (error) {
      toast.error(`Error adding students: ${error.message}`);
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <div className="container mx-auto p-4 bg-gray-100 min-h-screen">
      <Toaster position="top-right" />
      <h1 className="text-3xl font-bold mb-6 text-center text-indigo-800">
        Add Students
      </h1>
      <div className="max-w-4xl mx-auto bg-white p-6 rounded-lg shadow-md">
        <form onSubmit={handleSubmit} className="mb-6">
          <div className="mb-4">
            <label
              htmlFor="file"
              className="block text-gray-700 text-sm font-bold mb-2"
            >
              Upload Excel File:
            </label>
            <div className="relative border-2 border-gray-300 border-dashed rounded-md p-6 mt-1">
              <input
                type="file"
                id="file"
                ref={fileInputRef}
                onChange={handleFileChange}
                className="absolute inset-0 w-full h-full opacity-0 cursor-pointer"
                accept=".xlsx, .xls"
                required
              />
              <div className="text-center">
                <Upload className="mx-auto h-12 w-12 text-gray-400" />
                <p className="mt-1 text-sm text-gray-600">
                  {fileName || "Drop your Excel file here, or click to select"}
                </p>
              </div>
            </div>
          </div>
          {fileName && (
            <div className="flex items-center mt-2 text-sm text-gray-600">
              <File className="h-4 w-4 mr-2" />
              <span className="truncate">{fileName}</span>
              <button
                type="button"
                onClick={handleRemoveFile}
                className="ml-2 text-red-500 hover:text-red-700"
              >
                <X className="h-4 w-4" />
              </button>
            </div>
          )}
          <button
            type="submit"
            className="w-full mt-4 bg-indigo-500 hover:bg-indigo-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline transition duration-300 ease-in-out"
            disabled={isSubmitting || !file}
          >
            {isSubmitting ? "Adding students..." : "Upload Students"}
          </button>
        </form>
        
        {previewData && (
          <div className="mt-6">
            <h2 className="text-xl font-semibold mb-2">File Preview (First 20 rows)</h2>
            <div className="overflow-x-auto">
              <OutTable
                data={previewData.rows}
                columns={previewData.cols}
                tableClassName="min-w-full divide-y divide-gray-200"
                tableHeaderRowClass="bg-gray-50"
                tableHeaderCellClass="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                tableBodyRowClass="bg-white divide-y divide-gray-200"
                tableBodyCellClass="px-6 py-4 whitespace-nowrap text-sm text-gray-500"
              />
            </div>
          </div>
        )}
      </div>
    </div>
  );
};

export default UploadStudent;