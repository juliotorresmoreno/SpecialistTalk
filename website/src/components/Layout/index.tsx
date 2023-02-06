import React, { useContext } from 'react'
import NavBar from '../NavBar'
import GuessNavBar from '../GuessNavBar'
import Container from 'react-bootstrap/Container'
import Social from '../Social'
import Aside from '../Aside'
import { Messages } from '../Social/Messages'
import SocialContext from '../../contexts/SocialContext'
import withSession from '../../hoc/withSession'
import { useAppSelector } from '../../store/hooks'

type LayoutProps = {} & React.PropsWithChildren

type _LayoutProps = {} & LayoutProps

const _Layout: React.FC<_LayoutProps> = ({ children }) => {
  const session = useAppSelector((state) => state.auth.session)
  const { activeChat } = useContext(SocialContext)

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
          {activeChat ? <Messages /> : <div>{children}</div>}
          <Aside />
        </div>
      </Container>
    </>
  )
}

const Layout = withSession<LayoutProps>(_Layout)

export default Layout
