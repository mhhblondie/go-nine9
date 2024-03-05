import React, { useEffect, useState } from 'react';
import { getCookie, getJWT } from '../../AuthContext/AuthContext';
import { Link, useNavigate } from 'react-router-dom';

const AdminDashboard = () => {
  const [salons, setSalons] = useState([]);
  const [myToken, setMyToken] = useState('');
  const navigate = useNavigate();

  const getSalons = async (token) => {
    try {
      const response = await fetch('http://localhost:8097/api/admin/salons', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
        mode: 'cors',
      });

      if (!response.ok) {
        const errorResponse = await response.json();
        throw new Error(errorResponse.message || 'Échec de la requête');
      }

      const data = await response.json();
      setSalons(data);
    } catch (err) {
      console.error(err.message || 'Une erreur inattendue s\'est produite');
    }
  };


  const handleDeleteSalon = async (salonID) =>{
    try {
      const response = await fetch(`http://localhost:8097/api/management/salons/${salonID}`, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${myToken}`
        },
        mode: 'cors',
      });
  
      if (!response.ok) {
        const errorResponse = await response.json();
        throw new Error(errorResponse.message || 'Échec de la requête d\'inscription');
      }else{
        getSalons(myToken)
      }
  
    } catch (err) {
      throw new Error(err.message || 'Une erreur inattendue s\'est produite');
    }
  }

  useEffect(() => {
    const cookie = getCookie('authToken');
    const token = getJWT(cookie);
    const salonID = token.salonID;
    if (salonID !== null) {
      getSalons(cookie);
      setMyToken(cookie);
    } else {
      navigate('/login');
    }
  }, []);

  return (
    <div>
      <h1>Admin Dashboard</h1>
      <Link to={`/administrateur/createSalon`}> <button>Nouveau Salon</button></Link>
      {salons.length > 0 && (
        <table>
          <thead>
            <tr>
              <th>Nom</th>
              <th>Adresse</th>
              <th>Téléphone</th>
              <th>Description</th>
              <th>Action</th>
            </tr>
          </thead>
          <tbody>
            {salons.map((salon) => (
              <tr key={salon.ID}>
                <td>{salon.Name}</td>
                <td>{salon.Address}</td>
                <td>{salon.Phone}</td>
                <td>{salon.Description}</td>
                <td>
                  <button onClick={() => handleDeleteSalon(salon.ID)}>Supprimer</button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
      {salons.length === 0 && (
        <h1>
          Vous n'avez pas encore de salon.{' '}
          <Link to="/admin/create">Cliquez ici pour en créer un.</Link>
        </h1>
      )}
    </div>
  );
};

export default AdminDashboard;
