import { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
  ResponsiveContainer,
  PieChart,
  Pie,
  Cell,
} from "recharts";

const AdminSessionDashboard = () => {
  const [sessionData, setSessionData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const { sessionName } = useParams();

  useEffect(() => {
    fetch(`http://localhost:8000/session/details/${sessionName}`)
      .then((response) => {
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        return response.json();
      })
      .then((data) => {
        setSessionData(data.session);
        setLoading(false);
      })
      .catch((error) => {
        setError("Failed to fetch session data");
        setLoading(false);
      });
  }, [sessionName]);

  if (loading) return <div className="text-center mt-8">Loading...</div>;
  if (error)
    return <div className="text-center mt-8 text-red-500">{error}</div>;
  if (!sessionData)
    return <div className="text-center mt-8">No data available</div>;

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

  const COLORS = ["#0088FE", "#00C49F", "#FFBB28", "#FF8042", "#8884D8"];

  const coursesData = sessionData.Courses.map((course, index) => ({
    name: course.name,
    totalSeats: course.seats,
    seatsFilled: course.seats_filled,
    seatsAvailable: course.seats - course.seats_filled,
    color: COLORS[index % COLORS.length],
  }));

  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-3xl font-bold mb-8">{sessionData.name} Dashboard</h1>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
        <div className="bg-white shadow-md rounded-lg p-6">
          <h2 className="text-xl font-semibold mb-4">Session Information</h2>
          <p>
            <strong>Type:</strong> {sessionData.session_type}
          </p>
          <p>
            <strong>Status:</strong>
            <span
              className={`ml-2 inline-block px-2 py-1 rounded-full text-white text-sm ${getStatusColor(sessionData.status)}`}
            >
              {sessionData.status}
            </span>
          </p>
          <p>
            <strong>Created:</strong>{" "}
            {new Date(sessionData.CreatedAt).toLocaleString()}
          </p>
          <p>
            <strong>Last Updated:</strong>{" "}
            {new Date(sessionData.UpdatedAt).toLocaleString()}
          </p>
        </div>

        <div className="bg-white shadow-md rounded-lg p-6">
          <h2 className="text-xl font-semibold mb-4">Student Enrollment</h2>
          <div className="flex items-center justify-center h-full">
            <ResponsiveContainer width="100%" height={200}>
              <PieChart>
                <Pie
                  data={[
                    { name: "Applied", value: sessionData.applied_students },
                    {
                      name: "Remaining",
                      value:
                        sessionData.total_students -
                        sessionData.applied_students,
                    },
                  ]}
                  cx="50%"
                  cy="50%"
                  innerRadius={60}
                  outerRadius={80}
                  fill="#8884d8"
                  dataKey="value"
                  label={({ name, percent }) =>
                    `${name} ${(percent * 100).toFixed(0)}%`
                  }
                >
                  {[
                    <Cell key="cell-0" fill="#0088FE" />,
                    <Cell key="cell-1" fill="#00C49F" />,
                  ]}
                </Pie>
                <Tooltip />
              </PieChart>
            </ResponsiveContainer>
          </div>
          <p className="text-center mt-4">
            <strong>{sessionData.applied_students}</strong> out of{" "}
            <strong>{sessionData.total_students}</strong> students applied
          </p>
        </div>
      </div>

      <div className="bg-white shadow-md rounded-lg p-6 mb-8">
        <h2 className="text-xl font-semibold mb-4">Course Seat Allocation</h2>
        <ResponsiveContainer width="100%" height={300}>
          <BarChart data={coursesData}>
            <CartesianGrid strokeDasharray="3 3" />
            <XAxis dataKey="name" />
            <YAxis />
            <Tooltip />
            <Legend />
            <Bar
              dataKey="seatsFilled"
              stackId="a"
              fill="#8884d8"
              name="Seats Filled"
            />
            <Bar
              dataKey="seatsAvailable"
              stackId="a"
              fill="#82ca9d"
              name="Seats Available"
            />
          </BarChart>
        </ResponsiveContainer>
      </div>

      <div className="bg-white shadow-md rounded-lg p-6">
        <h2 className="text-xl font-semibold mb-4">Course Details</h2>
        <div className="overflow-x-auto">
          <table className="min-w-full table-auto">
            <thead>
              <tr className="bg-gray-200">
                <th className="px-4 py-2">Name</th>
                <th className="px-4 py-2">Code</th>
                <th className="px-4 py-2">Total Seats</th>
                <th className="px-4 py-2">Seats Filled</th>
                <th className="px-4 py-2">Seats Available</th>
                <th className="px-4 py-2">Fill Rate</th>
                <th className="px-4 py-2">Department ID</th>
              </tr>
            </thead>
            <tbody>
              {sessionData.Courses.map((course) => (
                <tr key={course.ID} className="border-b">
                  <td className="px-4 py-2">{course.name}</td>
                  <td className="px-4 py-2">{course.code}</td>
                  <td className="px-4 py-2">{course.seats}</td>
                  <td className="px-4 py-2">{course.seats_filled}</td>
                  <td className="px-4 py-2">
                    {course.seats - course.seats_filled}
                  </td>
                  <td className="px-4 py-2">
                    {((course.seats_filled / course.seats) * 100).toFixed(2)}%
                  </td>
                  <td className="px-4 py-2">{course.department_id}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
};

export default AdminSessionDashboard;
