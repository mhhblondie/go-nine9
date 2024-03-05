import React, { useContext, useState } from "react";
import "./navBar.css";
import { AuthContext } from "../../AuthContext/AuthContext";
import { User } from "lucide-react";
import { Link } from "react-router-dom";
import ModalUpdateUser from "../modalUpdateUser/ModalUpdateUser";
import ModalUpdatePassword from "../modalUpdatePassword/ModalUpdatePassword";

const Navbar = () => {
  const { isAuthenticated, logout } = useContext(AuthContext);
  const [menuVisible, setMenuVisible] = useState(false);
  const [showProfileModal, setShowProfileModal] = useState(false);
  const [showPassModal, setShowPassModal] = useState(false);

  const authTokenExists = !!document.cookie.includes('authToken');

  const handleLogout = () => {
    logout();
  };
  
  const toggleMenu = () => {
    setMenuVisible(!menuVisible);
  };

  return (
    <div className="navbar">
      <div className="logo">GONINE</div>
      {(isAuthenticated || authTokenExists) ? (
        <div className="user-section">
          <div className="user-icon" onClick={toggleMenu}>
            <User />
            <i className="fa fa-user" aria-hidden="true"></i>
          </div>
          {menuVisible && (
            <div className="dropdown-menu">
              <ul>
                <li onClick={() => setShowProfileModal(true)}>Modifier mon profil</li>
                <li onClick={() => setShowPassModal(true)}>Modifier mon mot de passe</li>
                <button className='btn-primary' onClick={handleLogout}>Se déconnecter</button>
              </ul>
            </div>
          )}
        </div>
      ) : (
        <div className="button">
          <Link to="/login"><button className='btn-primary'>Se connecter</button></Link>
        </div>
      )}
      <ModalUpdateUser show={showProfileModal} onClose={() => setShowProfileModal(false)} />
      <ModalUpdatePassword show={showPassModal} onClose={() => setShowPassModal(false)} />
    </div>
  );
};

export default Navbar;
