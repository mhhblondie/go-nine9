import React, {useEffect, useState} from 'react'
import { getCookie, getJWT } from '../../AuthContext/AuthContext';
import { Link, useNavigate } from "react-router-dom";


const Dashboard = () => {
  const [salon, setSalon] = useState();
  const [MyToken, setMyToken] = useState();
  const navigate = useNavigate();

  const getSalon = async (token) =>{
    try {
      const response = await fetch('http://localhost:8097/api/management/salons', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`
        },
        mode: 'cors',
      });
  
      if (!response.ok) {
        const errorResponse = await response.json();
        throw new Error(errorResponse.message || 'Échec de la requête d\'inscription');
      }

      const goodResponse = await response.json();
      console.log(goodResponse[0])
      setSalon(goodResponse[0])
  
    } catch (err) {
      throw new Error(err.message || 'Une erreur inattendue s\'est produite');
    }
  }

  const handleDeleteSalon = async () =>{
    try {
      const response = await fetch(`http://localhost:8097/api/management/salons/${salon.ID}`, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${MyToken}`
        },
        mode: 'cors',
      });
  
      if (!response.ok) {
        const errorResponse = await response.json();
        throw new Error(errorResponse.message || 'Échec de la requête d\'inscription');
      }else{
        navigate('/admin/create')
      }
  
    } catch (err) {
      throw new Error(err.message || 'Une erreur inattendue s\'est produite');
    }
  }


  const handleDeleteStaff = async (userID) =>{
    try {
      const response = await fetch(`http://localhost:8097/api/management/salons/staff/${userID}`, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${MyToken}`
        },
        mode: 'cors',
      });
  
      if (!response.ok) {
        const errorResponse = await response.json();
        throw new Error(errorResponse.message || 'Échec de la requête d\'inscription');
      }else{
        getSalon(MyToken)
      }
  
    } catch (err) {
      throw new Error(err.message || 'Une erreur inattendue s\'est produite');
    }
  }

  useEffect(() => {
    const cookie = getCookie('authToken');
    const token = getJWT(cookie)
    const salonID = token.salonID
    if(salonID !== null){
       getSalon(cookie)
       setMyToken(cookie)
    }else{
      navigate('/admin/create')
    }
  }, []);


  return (
    <div>
   <h1> Dashboard de mon salon </h1>
   {salon && <Link to="/admin/modify" state={{ state: salon }} ><button> Modifier mon salon </button></Link>}
   <Link to="/admin/addStaff"><button> Ajouter des salariés </button></Link>
   <button onClick={handleDeleteSalon}> Supprimer mon salon </button>
   {salon ?
   <>
    <h2>{salon.Name}</h2>
    <p>{salon.Description}</p>
    <p>adresse de mon salon : {salon.Address}</p>
    <p>Téléphone: {salon.Phone}</p>
    <h2>L'équipe</h2>
    {salon.User.map((user)=>(
      <>
      {user.Roles === "staff" &&   <button onClick={()=>{handleDeleteStaff(user.ID)}}>Supprimer</button>}
    
      <h2>{user.Firstname}  {user.Lastname} {user.Roles === "manager" && "(moi)"}</h2>
      <p>{user.Email} </p>
      <p>{user.Roles} </p>
      </>
    ))}

    </>  
    
   : <h1> Vous n'avez pas encore de salon snif <Link to="/admin/create"> Cliquez ici pour en créer un </Link></h1> }
      
    </div>
  )
}

export default Dashboard
