import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import ErrorPage from './Pages/ErrorPage';
import Home from './Pages/HomePage';
import AboutPage from './Pages/AboutPage';
import BlogsPage from './Pages/BlogsPage';
import Blog from './Pages/Blog';
import TutorialsPage from './Pages/TutorialsPage';
import Tutorial from './Pages/Tutorial';
import Admin from './Pages/Admin';
import Signup from './Pages/Signup';
import Login from './Pages/Login';
import SearchPage from './Pages/SearchPage';
import Auth from './Pages/Auth';

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <ErrorPage />,
    children: [
      {
        index: true,
        element: <Home />
      },
      {
        path: "/about_me",
        element: <AboutPage />
      },
      {
        path: "/blogs",
        element: <BlogsPage />
      },
      {
        path: "/blogs/:blog_id",
        element: <Blog />
      },
      {
        path: "/tutorials",
        element: <TutorialsPage />
      },
      {
        path: "/tutorials/:tutorial_id",
        element: <Tutorial />
      },
      {
        path: "/search/:query",
        element: <SearchPage />
      },
    ]
  },
  {
    path: "/auth",
    element: <Auth />,
    errorElement: <ErrorPage />,
    children: [
      {
        path: "login",
        element: <Login />
      },
      {
        path: "signup",
        element: <Signup />
      },
    ]
  },
  {
    path: "/admin",
    element: <Admin />,
    errorElement: <ErrorPage />,
    children: [{}]
  },
])

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
