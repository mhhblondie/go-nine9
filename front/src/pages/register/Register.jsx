// eslint-disable-next-line no-unused-vars
import React, { useContext, useEffect, useState } from 'react';
import { AuthContext, setCookie } from '../../AuthContext/AuthContext';
import { useNavigate } from 'react-router-dom';
import { Link } from 'react-router-dom';
import './register.css'
import { XCircle } from 'lucide-react';

function Register() {
  const [firstname, setFirstname] = useState('');
  const [lastname, setLastname] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [role, setRole] = useState("users");
  const [error, setError] = useState(null);
  const [rows, setRows] = useState([{ lastname: '', firstname: '', email: '' }]);
  const [salonInfo, setSalonInfo] = useState({ name: '', address: '', phone: '' ,description:''});
  
  const handleInputChange = (index, event) => {
    const values = [...rows];
    values[index][event.target.name] = event.target.value;
    setRows(values);
  };

  const handleInputChangeSalon = (event) => {
    setSalonInfo({
      ...salonInfo,
      [event.target.name]: event.target.value
  });
  };

  const increment = () => {
    setRows([...rows, { lastname: '', firstname: '', email: '' }]);
  };

  const decrement = (index) => {
      const values = [...rows];
      values.splice(index, 1);
      setRows(values);
  };

  const navigate = useNavigate();
  const { register, isAuthenticated } = useContext(AuthContext);

  useEffect(() => {

    if (isAuthenticated) {
      navigate('/');
    }
  }, [isAuthenticated, navigate, role]);

  const openPopup = (event) => {
    event.preventDefault();
    document.querySelector('.container-popup').style.display = 'flex';
  }

  const closePopup = () => {
    const popupContainer = document.querySelector('.container-popup');
    popupContainer.style.display = 'none';

    const inputs = popupContainer.querySelectorAll('input');
    inputs.forEach(input => {
        input.value = '';
    });

    const textareas = popupContainer.querySelectorAll('textarea');
    textareas.forEach(textarea => {
        textarea.value = '';
    });

    setSalonInfo({});
    setRows([{}]);
    setSalonInfo({});
};

  const registerPro = () => {
    const body = {
      firstname:firstname,
      Lastname:lastname,
      Email:email,
      Password:password,
      Roles:role,
      SalonInfo:salonInfo,
      Users: rows
  };

  fetch('http://localhost:8097/auth/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(body)
    })
    .then(response => response.json())
    .then(data => {
        setCookie('authToken', data.jwt, 1);
    });
    closePopup();
    navigate('/');
  }

  async function handleSubmit(event) {
    event.preventDefault();
    if (!firstname || !lastname || !email || !password || !role) {
      setError('Veuillez remplir tous les champs.');
      return;
    }
    try {
      await register(firstname, lastname, email, password, role);
      navigate('/');
    } catch (err) {
      setError(err.message);
    }
  }

  

  return (
    <div className="container-page">
      <div className="container-popup">
        <div className="card-popup">
          <XCircle className='close' onClick={closePopup}/>
          <h1>Créer votre établissement</h1>
              <div className="salon-info"> 
                <div className="top">
                <input
                      type="text"
                      name="name"
                      placeholder="Nom de l'établissement"
                      value={salonInfo.name}
                      onChange={handleInputChangeSalon}
                  />
                  <input
                      type="text"
                      name="address"
                      placeholder="Adresse"
                      value={salonInfo.address}
                      onChange={handleInputChangeSalon}
                  />
                  <input
                      type="text"
                      name="phone"
                      placeholder="Phone"
                      value={salonInfo.phone}
                      onChange={handleInputChangeSalon}
                  />
                </div>
                <textarea 
                  id="" 
                  cols="30" 
                  rows="10" 
                  name="description"
                  value={salonInfo.description} 
                  placeholder="Description"
                  onChange={handleInputChangeSalon}
                >
                </textarea>
              </div>
              <h2>Ajouté des employés</h2>
              {rows.map((row, index) => (
            <div className="row" key={index}>
                    <input
                        name="lastname"
                        id={`lastnameHairdressingSalon${index}`}
                        type="text"
                        placeholder='Nom'
                        value={row.lastname}
                        onChange={event => handleInputChange(index, event)}
                    />
                    <input
                        name="firstname"
                        id={`firstnameHairdressingSalon${index}`}
                        type="text"
                        placeholder='Prénom'
                        value={row.firstname}
                        onChange={event => handleInputChange(index, event)}
                    />
                    <input
                        name="email"
                        id={`emailHairdressingSalon${index}`}
                        type="text"
                        placeholder='Email'
                        value={row.email}
                        onChange={event => handleInputChange(index, event)}
                    />
                <button onClick={() => increment()}>+</button>
                {rows.length > 1 && (
                    <button onClick={() => decrement(index)}>-</button>
                )}
            </div>
        ))}
              <button className="btn-primary" onClick={registerPro}>Créer mon établissement</button>
        </div>
      </div>
      <div className="left">
        <div className="register-container">
          <h1 className="register-title">Nouveau sur Gonine ?</h1>
          <form action="" className="container">
            <div className="top-info">
              <label htmlFor="lastname">
                Nom *
                <input
                  type="text"
                  placeholder="Nom"
                  name="lastname"
                  className="form-control"
                  value={lastname}
                  onChange={(e) => setLastname(e.target.value)}
                />
              </label>
              <label htmlFor="firstname">
                Prénom *
                <input
                  type="text"
                  placeholder="Prénom"
                  name="firstname"
                  className="form-control"
                  value={firstname}
                  onChange={(e) => setFirstname(e.target.value)}
                />
              </label>
            </div>
            <label htmlFor="email">
              Email *
              <input
                type="text"
                placeholder="Email"
                name="email"
                className="form-control"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
              />
            </label>
            <label htmlFor="password">
              Mot de passe *
              <input
                type="password"
                placeholder="Mot de passe"
                className="form-control"
                name="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
            </label>
            
            <div className="role-container">
              <h2>Êtes-vous un professionnel</h2>
              <div className="role-content">
              <label htmlFor="dewey" className='roles'>
                  <input 
                      type="radio" 
                      id="manager"
                      name="user"
                      value="manager"
                      onChange={(e)=>{setRole(e.target.value)}}
                      checked={role === 'manager'}
                  />
                  Oui
              </label>
              <label htmlFor="huey" className='roles'>
                  <input 
                      type="radio"
                      id="users"
                      name="user"
                      value="users"
                      onChange={(e)=>{setRole(e.target.value)}}
                      checked={role === 'users'}
                  />
                  Non
              </label>
              </div>
            </div>
            {role === "manager" ? (
                <button className="btn-primary" onClick={(event) => openPopup(event)}>Créer mon compte</button>
            ):(
                <button className="btn-primary" onClick={(event) => handleSubmit(event)}>Créer mon compte</button>
              )
            }
            {error && <p className="error">{error}</p>}
          </form>
          <span className="separation">ou</span>
            <div className="register">
              <h2>Vous avez déjà utilisé Gonine?</h2>
              <Link to="/login" className='btn-secondary'>Créer mon compte</Link>
          </div>
        </div>
      </div>
      <div className="right">
        <img src="https://res.cloudinary.com/planity/image/upload/q_auto,f_auto/v1701340648/portail/illustrations/LOGIN/2023/3.jpg" alt="" />
      </div>
    </div>
  );
}

export default Register;