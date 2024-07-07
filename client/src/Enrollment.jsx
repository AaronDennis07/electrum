import { useState, useEffect } from "react";
import { useParams, Navigate } from "react-router-dom";
import toast, { Toaster } from "react-hot-toast";
import { useUser } from "./UserContext"; // Import the useUser hook
import { useAuth } from "./AuthContext";

const EnrollmentPeriodCourses = () => {
  const { sessionName } = useParams();
  const { user } = useAuth(); // Use the user.userId from context
  const [courses, setCourses] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [enrollingCourse, setEnrollingCourse] = useState(null);
  const [enrolled, setEnrolled] = useState(null);

  useEffect(() => {
    if (user.userId) {
      fetchCourses();
    }

  }, [user.userId, sessionName]);
  useEffect(() => {
    checkEnrollmentStatus();
  }, [enrolled]);

  const fetchCourses = async () => {
    setLoading(true);
    try {
      const response = await fetch(
        `http://localhost:8000/session/${sessionName}`,
      );
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }
      const data = await response.json();
      setCourses(
        data.courses.map((course) => ({
          ...course,
          availableSeats: course.Seats,
        })),
      );
      setLoading(false);
    } catch (error) {
      setError("Failed to fetch courses");
      setLoading(false);
      toast.error("Failed to load courses. Please try again later.");
    }
  };
  const checkEnrollmentStatus = async () => {
    try {
      const response = await fetch(
        `http://localhost:8000/session/${sessionName}/checkenrollment/${user.userId}`,
      );
      if (!response.ok) {
        throw new Error("Failed to check enrollment status");
      }
      const data = await response.json();
      if(data.enrolled){
        setEnrolled(data.coursecode)
      }
    } catch (error) {
      setError("Failed to check enrollment status");
      toast.error("Failed to check enrollment status.");
    }
};

  useEffect(() => {
    if (courses.length > 0) {
      const ws = new WebSocket(`ws://localhost:8000/session/ws/${sessionName}`);

      ws.onopen = () => {
        console.log("WebSocket Connected");
      };

      ws.onmessage = (event) => {
        const data = JSON.parse(event.data);
        setCourses((prevCourses) =>
          prevCourses.map((course) => ({
            ...course,
            availableSeats: parseInt(
              data[course.Code] || course.availableSeats,
            ),
          })),
        );
      };

      ws.onclose = () => {
        console.log("WebSocket Disconnected");
      };

      return () => {
        ws.close();
      };
    }
  }, [courses.length, sessionName]);

  const handleEnroll = async (courseCode) => {
    setEnrollingCourse(courseCode);
    setEnrolled(courseCode)
    try {
      const response = await fetch(
        `http://localhost:8000/session/${sessionName}/enroll`,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            id: user.userId,
            course: courseCode,
          }),
        },
      );

      if (!response.ok) {
        const data = await response.json();
        throw new Error(data.message);
      }

      toast.success(`Successfully enrolled in course ${courseCode}`);
    } catch (error) {
      toast.error(` ${error.message}`);
    } finally {
      setEnrollingCourse(null);
    }
  };

  if (!user.userId) {
    return <Navigate to="/login" />;
  }

  if (loading) return <div className="text-center mt-8">Loading...</div>;
  if (error)
    return <div className="text-center mt-8 text-red-500">{error}</div>;

  return (
    <div className="container mx-auto p-4 bg-gray-100">
      <Toaster position="top-right" />
      <h1 className="text-3xl font-bold mb-6 text-center text-indigo-800">
        Course Enrollment for {sessionName}
      </h1>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {courses.map((course) => (
          <div
            key={course.Id}
            className="bg-white rounded-lg shadow-lg overflow-hidden hover:shadow-xl transition-shadow duration-300 border border-indigo-100"
          >
            <div className="p-6">
              <h2 className="text-xl font-semibold mb-2 text-indigo-700">
                {course.Name}
              </h2>
              <p className="text-gray-600 mb-1">Course Code: {course.Code}</p>
              <p className="text-gray-600 mb-4">
                Department: {course.Department}
              </p>
              <div className="mb-4">
                <div className="flex justify-between items-center mb-1">
                  <span className="text-sm font-medium text-gray-700">
                    Available Seats:
                  </span>
                  <span className="text-sm font-medium text-indigo-600">
                    {course.availableSeats} / {course.Seats}
                  </span>
                </div>
                <div className="w-full bg-gray-200 rounded-full h-2.5">
                  <div
                    className="bg-indigo-600 h-2.5 rounded-full transition-all duration-500 ease-in-out"
                    style={{
                      width: `${(course.availableSeats / course.Seats) * 100}%`,
                    }}
                  ></div>
                </div>
              </div>
              <button
                className="w-full bg-indigo-500 hover:bg-indigo-700 text-white font-bold py-2 px-4 rounded transition duration-300 disabled:bg-gray-300 disabled:cursor-not-allowed"
                onClick={() => handleEnroll(course.Code)}
                disabled={
                  enrollingCourse === course.Code || course.availableSeats === 0 || enrolled !== null
                }
              >
                {enrollingCourse === course.Code
                  ? "Enrolling..."
                  : enrolled === course.Code
                    ? "Enrolled"
                    : course.availableSeats === 0
                      ? "Full"
                      : "Enroll"}
              </button>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default EnrollmentPeriodCourses;
