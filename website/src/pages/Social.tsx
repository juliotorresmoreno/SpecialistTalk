import React from 'react'
import Header from '../components/Header'

const SocialPage: React.FC = () => {
  const header = {
    title: 'Social',
    description: 'programa de super poderes',
  }

  return (
    <>
      <Header {...header} />
      <main>Recommends</main>
    </>
  )
}

export default SocialPage
