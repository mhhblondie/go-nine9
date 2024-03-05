import { useContext, useEffect, useState } from 'react';
import './login.css';
import { AuthContext } from '../../AuthContext/AuthContext';
import { Link, useNavigate } from 'react-router-dom';

function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState(null);
  const { login, isAuthenticated } = useContext(AuthContext);
  const navigate = useNavigate();

  useEffect(() => {
    const checkAuthentication = async () => {
      if (isAuthenticated) {
        navigate('/');
      }
    };

    checkAuthentication();
  }, [isAuthenticated, navigate]);

  async function handleSubmit(event) {
    event.preventDefault();
    if (!email || !password) {
      setError('Veuillez remplir tout les champs.');
      return;
    }
    try {
      await login(email, password);
      navigate('/');
    } catch (err) {
      setError(err.message);
    }
  }

  return (
    <div className="container-page">
      <div className="left">
        <div className="login-container">
          <h1 className="login-title">Vous avez déjà utilié Gonine ?</h1>
          <form action="" className="form-container">
            <label htmlFor="emailInput" className="email">
              Email *
              <input
                id="emailInput"
                type="text"
                placeholder="Email"
                name="email"
                className="form-control"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
              />
            </label>
            <label htmlFor="passwordInput" className='password'>
              Mot de passe *
              <input
                id='passwordInput'
                type="password"
                placeholder="Mot de passe"
                className="form-control"
                name="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
            </label>
            <button onClick={(event) => handleSubmit(event)} className='btn-primary'>Se connecter</button>
            {error && <p className="error">{error}</p>}
          </form>
          <span className="separation">ou</span>
          <div className="register">
            <h2>Nouveau sur Gonine?</h2>
            <Link to="/register" className='btn-secondary'>Créer mon compte</Link>
          </div>
        </div>
      </div>
      <div className="right">
        <img src="https://res.cloudinary.com/planity/image/upload/q_auto,f_auto/v1701340648/portail/illustrations/LOGIN/2023/3.jpg" alt="" />
      </div>
    </div>
  );
}

export default Login;
