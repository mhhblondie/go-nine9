import React, { useState } from 'react';
import './modalUpdateUSer.css';
import { getCookie, setCookie } from '../../AuthContext/AuthContext'; 

const ModalUpdateUser = ({ show, onClose}) => {
    const authToken = getCookie('authToken'); 
    const [formData, setFormData] = useState({
        lastname: "",
        firstname: "",
        email: "",
    });

    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData(prevState => ({
            ...prevState,
            [name]: value
        }));
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const response = await fetch("http://localhost:8097/api/me", {
                method: "PATCH",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": `Bearer ${authToken}`
                },
                mode: 'cors',
                body: JSON.stringify({
                    lastname: formData.lastname,
                    firstname: formData.firstname,
                    email: formData.email,
                })
            });
            if (response.ok) {
                const newToken = response.headers.get('Authorization');
                if (newToken) {
                    setCookie('authToken', newToken, 1);
                }
                onClose();
            } else {
                console.error("Échec de la modification du profil");
            }
        } catch (error) {
            console.error("Erreur lors de la modification du profil :", error);
        }
    };

    if (!show) {
        return null;
    }

    console.log(authToken);
    
  return (
    <div className="modal">
      <div className="modal-content">
        <span className="close" onClick={onClose}>&times;</span>
        <h2>Modifier mon profil</h2>
        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <label htmlFor="lastname">Nom :</label>
            <input type="text" id="lastname" name="lastname" value={formData.lastname} onChange={handleChange} />
          </div>
          <div className="form-group">
            <label htmlFor="firstname">Prénom :</label>
            <input type="text" id="firstname" name="firstname" value={formData.firstname} onChange={handleChange} />
          </div>
          <div className="form-group">
            <label htmlFor="email">Email :</label>
            <input type="email" id="email" name="email" value={formData.email} onChange={handleChange} />
          </div>
          <button type="submit">Valider</button>
        </form>
      </div>
    </div>
  );
}

export default ModalUpdateUser;
