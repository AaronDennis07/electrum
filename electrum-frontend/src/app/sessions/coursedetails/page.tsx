/**
 * v0 by Vercel.
 * @see https://v0.dev/t/iwOOu1otJER
 * Documentation: https://v0.dev/docs#integrating-generated-code-into-your-nextjs-app
 */
"use client"

import { JSX, SVGProps, useEffect, useState } from "react"
import { Button } from "@/components/ui/button"
import axios from "axios"

type Course = {
  Id: number
  Name: string
  Code: string
  Department: string
  Seats: number
}
export default function Page() {
  const [courses, setCourses] = useState<Course[]>([])
  const [showToast, setShowToast] = useState(false)

  useEffect(() => {
    axios.get("http://localhost:8000/session/sessiontry").then((response) => {
      setCourses(response.data.courses)
      console.log(response.data.courses)
      console.log(Array.isArray(response.data.courses))
      console.log(courses)
    }).catch((error) => {
      console.log(error)
    })
  }, [])

  const handleEnroll = (courseId: number) => {
    setCourses((prevCourses) => {
      return prevCourses?.map((course) => {
        if (course.Id === courseId && course.Seats > 0) {
          return { ...course, remainingSeats: course.Seats - 1, isEnrolled: true }
        } else if (course.Id !== courseId) {
          return { ...course, isEnrolled: false }
        }
        return course
      })
    })
    setShowToast(true)
    setTimeout(() => {
      setShowToast(false)
    }, 3000)
  }
  return (
    <section className="w-full py-12">
      <div className="container grid gap-6 md:gap-8 px-4 md:px-6">
        <div className="flex flex-col md:flex-row items-start md:items-center gap-4 md:gap-8">
          <div className="grid gap-1">
            <h1 className="text-2xl font-bold tracking-tight">Course Selection</h1>
            <p className="text-gray-500 dark:text-gray-400">Browse and enroll in available courses.</p>
          </div>
        </div>
        <ul className="grid sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-8">
          {courses && courses?.map((course) => (
        
            <li
              key={course.Id}
              className="grid gap-4 p-4 border rounded-lg shadow-sm hover:shadow-md transition-shadow"
            >
              <div className="grid gap-2">
                <div className="flex items-center gap-4">
                  <h3 className="font-semibold">{course.Name}</h3>
                  <span className="text-sm font-medium text-gray-500 dark:text-gray-400">{course.Code}</span>
                </div>
                <p className="text-sm leading-none text-gray-500 dark:text-gray-400">{course.Department}</p>
              </div>
              <div className="flex items-center justify-between">
                <div className="flex items-center gap-2 text-sm font-medium">
                  <UserIcon className="w-4 h-4 text-gray-500 dark:text-gray-400" />
                  <span>
                    {course.Seats}/{course.Seats}
                  </span>
                </div>
                <Button
                  size="sm"
                  variant="outline"
                  // disabled={course.seats === 0 || course.isEnrolled}
                  onClick={() => handleEnroll(course.Id)}
                >
                  {/* {course.seats === 0 ? "Full" : course.isEnrolled ? "Enrolled" : "Enroll"} */}
                  {course.Seats === 0 ? "Full" :  "Enroll"}

                </Button>
              </div>
            </li>
          ))}
        </ul>
      </div>
      {showToast && (
        <div className="fixed bottom-4 left-1/2 transform -translate-x-1/2 bg-green-500 text-white px-4 py-2 rounded-md shadow-lg">
          You have successfully enrolled in a course.
        </div>
      )}
    </section>
  )
}

function UserIcon(props: JSX.IntrinsicAttributes & SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2" />
      <circle cx="12" cy="7" r="4" />
    </svg>
  )
}