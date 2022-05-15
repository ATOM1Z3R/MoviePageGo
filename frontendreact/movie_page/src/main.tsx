import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import './index.css'
import MoviesPage from './pages/Movies';
import DirectorsPage from './pages/Directros';
import GenresPage from './pages/Genres';
import LoginPage from './pages/Login';
import SignUpPage from './pages/Signup';


ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<App />} >
          <Route path="movies" element={<MoviesPage />} />
          <Route path="directors" element={<DirectorsPage />} />
          <Route path="genres" element={<GenresPage />} />
          <Route path="login" element={<LoginPage />} />
          <Route path="signup" element={<SignUpPage />} />
        </Route>
      </Routes>
    </BrowserRouter>
  </React.StrictMode>
)
