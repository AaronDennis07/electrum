/**
* This code was generated by v0 by Vercel.
* @see https://v0.dev/t/aIwVcLsIcj9
* Documentation: https://v0.dev/docs#integrating-generated-code-into-your-nextjs-app
*/

/** Add fonts into your Next.js project:

import { Cormorant_Garamond } from 'next/font/google'
import { Archivo } from 'next/font/google'

cormorant_garamond({
  subsets: ['latin'],
  display: 'swap',
})

archivo({
  subsets: ['latin'],
  display: 'swap',
})

To read more about using these font, please visit the Next.js documentation:
- App Directory: https://nextjs.org/docs/app/building-your-application/optimizing/fonts
- Pages Directory: https://nextjs.org/docs/pages/building-your-application/optimizing/fonts
**/
import { Button } from "@/components/ui/button"
import { CardTitle, CardHeader, CardContent, Card } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"

export function Sessiondetails() {
  return (
    <div className="grid gap-8 p-6 md:p-8 lg:p-10">
      <div className="grid gap-4">
        <div className="flex items-center justify-between">
          <div className="grid gap-1">
            <h1 className="text-2xl font-bold">Session Details</h1>
            <p className="text-gray-500 dark:text-gray-400">Review the details of this session.</p>
          </div>
          <Button size="sm" variant="outline">
            <DownloadIcon className="h-4 w-4 mr-2" />
            Download as Excel
          </Button>
        </div>
        <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
          <Card>
            <CardHeader>
              <CardTitle>Session Name</CardTitle>
            </CardHeader>
            <CardContent>
              <p className="text-2xl font-bold">Summer Bootcamp 2023</p>
            </CardContent>
          </Card>
          <Card>
            <CardHeader>
              <CardTitle>Session Type</CardTitle>
            </CardHeader>
            <CardContent>
              <p className="text-2xl font-bold">Online</p>
            </CardContent>
          </Card>
          <Card>
            <CardHeader>
              <CardTitle>Enrollment Closes In</CardTitle>
            </CardHeader>
            <CardContent>
              <p className="text-2xl font-bold">2 days, 4 hours</p>
            </CardContent>
          </Card>
        </div>
      </div>
      <div className="grid gap-6">
        <div className="grid gap-2">
          <div className="flex items-center justify-between">
            <h2 className="text-xl font-bold">Courses</h2>
            <Badge className="bg-gray-100 px-3 py-1 text-sm font-medium text-gray-900 dark:bg-gray-800 dark:text-gray-50">
              12 courses
            </Badge>
          </div>
          <div className="grid gap-4">
            <Card>
              <CardHeader>
                <CardTitle>Introduction to Web Development</CardTitle>
              </CardHeader>
              <CardContent className="grid grid-cols-[1fr_auto] items-center gap-4">
                <div className="grid gap-1">
                  <div className="flex items-center gap-2">
                    <div className="h-2 w-full rounded-full bg-gray-200 dark:bg-gray-800">
                      <div
                        className="h-2 rounded-full bg-gray-900 dark:bg-gray-50"
                        style={{
                          width: "75%",
                        }}
                      />
                    </div>
                    <span className="text-sm font-medium">75/100</span>
                  </div>
                  <div className="flex items-center gap-2 text-sm text-gray-500 dark:text-gray-400">
                    <BookIcon className="h-4 w-4" />
                    <span>WEB101</span>
                  </div>
                </div>
                <Badge className="bg-gray-100 px-3 py-1 text-sm font-medium text-gray-900 dark:bg-gray-800 dark:text-gray-50">
                  75% registered
                </Badge>
              </CardContent>
            </Card>
            <Card>
              <CardHeader>
                <CardTitle>Advanced JavaScript</CardTitle>
              </CardHeader>
              <CardContent className="grid grid-cols-[1fr_auto] items-center gap-4">
                <div className="grid gap-1">
                  <div className="flex items-center gap-2">
                    <div className="h-2 w-full rounded-full bg-gray-200 dark:bg-gray-800">
                      <div
                        className="h-2 rounded-full bg-gray-900 dark:bg-gray-50"
                        style={{
                          width: "90%",
                        }}
                      />
                    </div>
                    <span className="text-sm font-medium">90/100</span>
                  </div>
                  <div className="flex items-center gap-2 text-sm text-gray-500 dark:text-gray-400">
                    <BookIcon className="h-4 w-4" />
                    <span>JS201</span>
                  </div>
                </div>
                <Badge className="bg-gray-100 px-3 py-1 text-sm font-medium text-gray-900 dark:bg-gray-800 dark:text-gray-50">
                  90% registered
                </Badge>
              </CardContent>
            </Card>
            <Card>
              <CardHeader>
                <CardTitle>Introduction to React</CardTitle>
              </CardHeader>
              <CardContent className="grid grid-cols-[1fr_auto] items-center gap-4">
                <div className="grid gap-1">
                  <div className="flex items-center gap-2">
                    <div className="h-2 w-full rounded-full bg-gray-200 dark:bg-gray-800">
                      <div
                        className="h-2 rounded-full bg-gray-900 dark:bg-gray-50"
                        style={{
                          width: "60%",
                        }}
                      />
                    </div>
                    <span className="text-sm font-medium">60/100</span>
                  </div>
                  <div className="flex items-center gap-2 text-sm text-gray-500 dark:text-gray-400">
                    <BookIcon className="h-4 w-4" />
                    <span>REACT101</span>
                  </div>
                </div>
                <Badge className="bg-gray-100 px-3 py-1 text-sm font-medium text-gray-900 dark:bg-gray-800 dark:text-gray-50">
                  60% registered
                </Badge>
              </CardContent>
            </Card>
          </div>
        </div>
      </div>
    </div>
  )
}

function BookIcon(props) {
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
      <path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1 0-5H20" />
    </svg>
  )
}


function DownloadIcon(props) {
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
      <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4" />
      <polyline points="7 10 12 15 17 10" />
      <line x1="12" x2="12" y1="15" y2="3" />
    </svg>
  )
}