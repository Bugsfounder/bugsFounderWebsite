import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import ErrorPage from './Pages/ErrorPage';
import Home from './Pages/HomePage';
import AboutPage from './Pages/AboutPage';
import BlogsPage from './Pages/BlogsPage';
import BlogPage from './Pages/BlogPage';
import TutorialsPage from './Pages/TutorialsPage';
import TutorialPage from './Pages/TutorialPage';
import Admin from './Pages/admin/AdminPage';
import SignupPage from './Pages/SignupPage';
import LoginPage from './Pages/LoginPage';
import SearchPage from './Pages/SearchPage';
import Admins from './Components/AdminComponents/Admins';
import Blogs from './Components/AdminComponents/Blogs';
import Tutorials from './Components/AdminComponents/Tutorials';
import Users from './Components/AdminComponents/Users';

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
        element: <BlogPage />
      },
      {
        path: "/tutorials",
        element: <TutorialsPage />
      },
      {
        path: "/tutorials/:tutorial_id",
        element: <TutorialPage />
      },
      {
        path: "/search/:query",
        element: <SearchPage />
      },

    ]
  },
  {
    path: "/auth",
    element: <App />,
    errorElement: <ErrorPage />,
    children: [
      {
        path: "login",
        element: <LoginPage />
      },
      {
        path: "signup",
        element: <SignupPage />
      },
    ]
  },
  {
    path: "/kubari/admin",
    element: <Admin />,
    errorElement: <ErrorPage />,
    children: [
      {
        path: "admins",
        element: <Admins />
      },
      {
        path: "blogs",
        element: <Blogs />
      },
      {
        path: "tutorials",
        element: <Tutorials />
      },
      {
        path: "users",
        element: <Users />
      },
    ]
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
