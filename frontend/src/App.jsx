import { RouterProvider } from 'react-router-dom'
import { router } from './routes'

function App() {
  return <RouterProvider router={router} />
}

export default App


// export default function App() {
//   return (
//    <div
//   className="min-h-screen bg-blue-500 flex items-center justify-center"
//   style={{ color: 'red' }}
// >
//   Tailwind is working ✅
// </div>

//   );
// }

