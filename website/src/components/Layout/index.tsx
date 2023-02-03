import React from 'react'
import NavBar from '../NavBar'
import GuessNavBar from '../GuessNavBar'
import Container from 'react-bootstrap/Container'
import { useAppSelector } from '../../store/hooks'
import Social from '../Social'
import Aside from '../Aside'

type LayoutProps = {} & React.PropsWithChildren

const Layout: React.FC<LayoutProps> = ({ children }) => {
  const session = useAppSelector((state) => state.auth.session)

  return (
    <>
      {session ? (
        <>
          <GuessNavBar />
          <Social />
        </>
      ) : (
        <NavBar />
      )}

      <Container id="container">
        <div className="content">
          <div>{children}</div>
          <Aside />
        </div>
      </Container>
    </>
  )
}

export default Layout
