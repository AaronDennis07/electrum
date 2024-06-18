/**
* This code was generated by v0 by Vercel.
* @see https://v0.dev/t/MuruVt6iXop
* Documentation: https://v0.dev/docs#integrating-generated-code-into-your-nextjs-app
*/
import { Button } from "@/components/ui/button"
import { CardContent, Card } from "@/components/ui/card"

export function Sessionlist() {
  return (
    <div className="w-full max-w-6xl mx-auto py-12 md:py-16 lg:py-20">
      <h1 className="text-3xl font-bold mb-8 md:text-4xl">Project Dashboard</h1>
      <div className="space-y-8">
        <section>
          <h2 className="text-xl font-semibold mb-4">Live Projects</h2>
          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            <Card className="w-[350px]">
              <div className="relative h-[200px] overflow-hidden rounded-t-lg">
                <img alt="Project cover" className="object-cover" fill src="/placeholder.svg" />
              </div>
              <CardContent className="space-y-4">
                <div>
                  <h2 className="text-xl font-semibold">Project Name</h2>
                  <p className="text-gray-500 dark:text-gray-400">Start date: June 14, 2024</p>
                </div>
                <div className="flex justify-between">
                  <Button className="flex items-center gap-1">
                    <ClockIcon className="h-4 w-4" />
                    Apply Before 5:00 PM
                  </Button>
                  <Button>Apply</Button>
                </div>
              </CardContent>
            </Card>
          </div>
        </section>
        <section>
          <h2 className="text-xl font-semibold mb-4">Upcoming Projects</h2>
          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            <Card className="w-[350px]">
              <div className="relative h-[200px] overflow-hidden rounded-t-lg">
                <img alt="Project cover" className="object-cover" fill src="/placeholder.svg" />
              </div>
              <CardContent className="space-y-4">
                <div>
                  <h2 className="text-xl font-semibold">Project Name</h2>
                  <p className="text-gray-500 dark:text-gray-400">Start date: July 1, 2024</p>
                </div>
                <div className="flex justify-between">
                  <Button variant="outline">More Info</Button>
                </div>
              </CardContent>
            </Card>
          </div>
        </section>
        <section>
          <h2 className="text-xl font-semibold mb-4">Completed Projects</h2>
          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            <Card className="w-[350px]">
              <div className="relative h-[200px] overflow-hidden rounded-t-lg">
                <img alt="Project cover" className="object-cover" fill src="/placeholder.svg" />
              </div>
              <CardContent className="space-y-4">
                <div>
                  <h2 className="text-xl font-semibold">Project Name</h2>
                  <p className="text-gray-500 dark:text-gray-400">Completed: May 30, 2024</p>
                </div>
                <div className="flex justify-between">
                  <Button variant="outline">View Details</Button>
                </div>
              </CardContent>
            </Card>
          </div>
        </section>
      </div>
    </div>
  )
}

function ClockIcon(props) {
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
      <circle cx="12" cy="12" r="10" />
      <polyline points="12 6 12 12 16 14" />
    </svg>
  )
}