import React, { useRef, useState } from 'react'
import Container from 'react-bootstrap/Container'
import Nav from 'react-bootstrap/Nav'
import Navbar from 'react-bootstrap/Navbar'
import { Link, useLocation } from 'react-router-dom'
import classnames from 'classnames'
import useOnScreen from '../hooks/useOnScreen'
import Modal from './Modal'
import SingInForm from './SingInForm'
import SingUpForm from './SingUpForm'
import { Session } from '../models/session'
import authSlice from '../features/auth'
import { useAppDispatch } from '../store/hooks'

type NavBarProps = {} & React.PropsWithChildren

const NavBar: React.FC<NavBarProps> = (props) => {
  const toggleBtnRef = useRef<HTMLButtonElement | null>()
  const isVisible = useOnScreen(toggleBtnRef)
  const [showModal, setShowModal] = useState<Link | null>(null)
  const location = useLocation()
  const dispatch = useAppDispatch()
  const links0 = [
    ['/forum', 'Foro'],
    //['/recommends', 'Recomendado'],
    ['/social', 'Social'],
  ]
  const links1: Link[] = [
    {
      title: 'Ingresar',
      form: (
        <SingInForm
          onSuccess={(session: Session) => {
            setShowModal(null)
            dispatch(authSlice.actions.setSession(session))
          }}
        />
      ),
    },
    {
      title: 'Registrate',
      form: <SingUpForm onSubmit={() => {}} errors={{}} />,
    },
  ]
  const toggleNav = () => {
    if (isVisible) toggleBtnRef.current?.click()
  }
  const renderLinks = ([link, title]: string[]) => {
    const classNames = classnames({
      'nav-link': true,
      active: location.pathname === link,
    })
    return (
      <Link key={link} className={classNames} to={link} onClick={toggleNav}>
        {title}
      </Link>
    )
  }
  type Link = {
    title: string
    form: React.ReactNode
  }

  const renderLinksAuth = (link: Link) => {
    return (
      <Link
        key={link.title}
        to=""
        className="nav-link"
        onClick={() => setShowModal(link)}
      >
        {link.title}
      </Link>
    )
  }
  return (
    <Navbar collapseOnSelect expand="lg" bg="purple" variant="dark">
      <Container>
        <Link className="navbar-brand" to="/">
          Home
        </Link>
        <Navbar.Toggle
          ref={toggleBtnRef as any}
          aria-controls="responsive-navbar-nav"
        />
        <Navbar.Collapse id="responsive-navbar-nav">
          <Nav className="me-auto">{links0.map(renderLinks)}</Nav>
          <Nav>{links1.map(renderLinksAuth)}</Nav>
        </Navbar.Collapse>
        <Modal
          show={showModal !== null}
          title={showModal?.title ?? ''}
          handleSave={() => {}}
          handleClose={() => setShowModal(null)}
        >
          {showModal?.form}
        </Modal>
      </Container>
    </Navbar>
  )
}

export default NavBar
