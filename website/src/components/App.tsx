import React from 'react'
import { Route, Routes } from 'react-router-dom'
import withSession from '../hoc/withSession'
import ForumPage from '../pages/Forum'
import HomePage from '../pages/Home'
import MessagesPage from '../pages/Messages'
import NotFoundPage from '../pages/NotFound'
import ProfilePage from '../pages/Profile'
import RecommendsPage from '../pages/Recommends'
import SignInPage from '../pages/SignIn'
import SignUpPage from '../pages/SignUp'
import { useAppSelector } from '../store/hooks'
import Layout from './Layout'
import Provider from './Provider'

const _App: React.FC = () => {
  const session = useAppSelector((state) => state.auth.session)

  return (
    <Provider>
      <Layout>
        <Routes>
          {session ? (
            <>
              <Route path="/" element={<HomePage />} />
              <Route path="/profile" element={<ProfilePage />} />
              <Route path="/chats/:id" element={<MessagesPage />} />

              <Route path="*" element={<NotFoundPage />} />
            </>
          ) : (
            <>
              <Route path="/" element={<HomePage />} />
              <Route path="/forum" element={<ForumPage />} />
              <Route path="/recommends" element={<RecommendsPage />} />

              <Route path="/sign-in" element={<SignInPage />} />
              <Route path="/sign-up" element={<SignUpPage />} />

              <Route path="*" element={<NotFoundPage />} />
            </>
          )}
        </Routes>
      </Layout>
    </Provider>
  )
}

const App = withSession(_App)

export default App
