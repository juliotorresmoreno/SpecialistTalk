import React from 'react'
import { Helmet } from 'react-helmet'

export type HeaderProps = {
  title: string
  description: string
}

const Header: React.FC<HeaderProps> = ({ title, description }) => {
  return (
    <>
      <header>
        <h1>{title}</h1>
        <Helmet>
          {title && <title>{title} | SpecialistTalk</title>}
          {description && <meta name="description" content={description} />}
        </Helmet>
      </header>
    </>
  )
}

export default Header
