import React from 'react'
import Header from '../components/Header'

const ProfilePage: React.FC = () => {
  const header = {
    title: 'Profile',
    description: 'programa de super poderes',
  }
  return (
    <>
      <Header {...header} />
      <main>
        <div
          style={{
            display: 'flex',
            flex: 1,
          }}
        >
          <div></div>
        </div>
      </main>
    </>
  )
}

export default ProfilePage
