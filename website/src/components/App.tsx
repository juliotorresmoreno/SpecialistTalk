import React from 'react'
import { Route, Routes } from 'react-router-dom'
import ForumPage from '../pages/Forum'
import HomePage from '../pages/Home'
import NotFoundPage from '../pages/NotFound'
import RecommendsPage from '../pages/Recommends'
import SignInPage from '../pages/SignIn'
import SignUpPage from '../pages/SignUp'
import SocialPage from '../pages/Social'
import Layout from './Layout'

const App: React.FC = () => {
  return (
    <Layout>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/forum" element={<ForumPage />} />
        <Route path="/recommends" element={<RecommendsPage />} />
        <Route path="/social" element={<SocialPage />} />
        <Route path="/sign-in" element={<SignInPage />} />
        <Route path="/sign-up" element={<SignUpPage />} />

        <Route path="*" element={<NotFoundPage />} />
      </Routes>
    </Layout>
  )
}

export default App
