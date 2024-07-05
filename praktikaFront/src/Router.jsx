import App from './App';
import { createBrowserRouter, RouterProvider } from "react-router-dom";

const router = createBrowserRouter([
    {
        path: "/",
        element: <App />,
    },
]);

function MyRouter() {
    return (
        <>
            <RouterProvider router={router} />
        </>
    );
}

export default MyRouter;