import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import AdminApp from './AdminApp';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import ErrorPage from './Pages/ErrorPage';
import Home from './Pages/MainWebPages/HomePage';
import AboutPage from './Pages/MainWebPages/AboutPage';
import BlogsPage from './Pages/MainWebPages/BlogsPage';
import BlogPage from './Pages/MainWebPages/BlogPage';
import TutorialsPage from './Pages/MainWebPages/TutorialsPage';
import TutorialPage from './Pages/MainWebPages/TutorialPage';
import SignupPage from './Pages/AuthPages/SignupPage';
import LoginPage from './Pages/AuthPages/LoginPage';
import SearchPage from './Pages/MainWebPages/SearchPage';
import AdminBlogsPage from './Pages/AdminPages/AdminBlogsPage';
import AdminTutorialsPage from './Pages/AdminPages/AdminTutorialsPage';
import AdminUsersPage from './Pages/AdminPages/AdminUsersPage';
import AdminAdminsPage from './Pages/AdminPages/AdminAdminsPage';
import AdminPage from './Pages/AdminPages/AdminPage';
import AdminEditBlogPage from './Pages/AdminPages/AdminEditBlogPage';
import AdminAddBlogPage from './Pages/AdminPages/AdminAddBlogPage';
import 'react-notifications/lib/notifications.css';
import AdminEditTutorialPage from './Pages/AdminPages/AdminEditTutorialPage';
import AdminAddTutorialPage from './Pages/AdminPages/AdminAddTutorialPage';
import AdminAddSubTutorialPage from './Pages/AdminPages/AdminAddSubTutorialPage';
import AdminGetAllSubTutorialPage from './Pages/AdminPages/AdminGetAllSubTutorialPage';
import AdminUpdateSubTutorialPage from './Pages/AdminPages/AdminUpdateSubTutorialPage';

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
        path: "/blogs/:blog_url",
        element: <BlogPage />
      },
      {
        path: "/tutorials",
        element: <TutorialsPage />
      },
      {
        path: "/tutorials/:tutorial_url/:sub_tutorial_url",
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
    element: <AdminApp />,
    errorElement: <ErrorPage />,
    children: [
      {
        path: "",
        element: <AdminPage />
      },
      {
        path: "admins",
        element: <AdminAdminsPage />
      },
      {
        path: "blogs",
        element: <AdminBlogsPage />
      },
      {
        path: "blogs/add",
        element: <AdminAddBlogPage />
      },
      {
        path: "blogs/edit/:blog_url",
        element: <AdminEditBlogPage />
      },
      {
        path: "tutorials",
        element: <AdminTutorialsPage />
      },
      {
        path: "tutorials/edit/:tutorial_url",
        element: <AdminEditTutorialPage />
      },
      {
        path: "tutorials/edit/sub_tutorial/add/:tutorial_url",
        element: <AdminAddSubTutorialPage />
      },
      {
        path: "tutorials/edit/sub_tutorial/:tutorial_url",
        element: <AdminGetAllSubTutorialPage />
      },
      {
        path: "tutorials/edit/sub_tutorial/:tutorial_url/:sub_tutorial_url",
        element: <AdminUpdateSubTutorialPage />
      },
      {
        path: "tutorials/add",
        element: <AdminAddTutorialPage />
      },
      {
        path: "users",
        element: <AdminUsersPage />
      },
    ]
  },
])

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  // React.StrictMode render components twice which is causing problems with fetching data and appending into a state it inserts duplicate data on second render
  // <React.StrictMode> 
  <RouterProvider router={router} />
  // </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
