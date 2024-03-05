import React, { useState } from 'react';
import { getCookie } from '../../AuthContext/AuthContext';

const ModalUpdatePassword = ({ show, onClose }) => {
    const authToken = getCookie('authToken');
    const [formData, setFormData] = useState({
        password: "",
})
    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData(prevState => ({
            ...prevState,
            [name]: value
        }));
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        try{
            const response = await fetch("http://localhost:8097/api/me/password", {
                method: "PATCH",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": `Bearer ${authToken}`
                },
                mode: 'cors',
                body: JSON.stringify({
                    password: formData.password
                })
            });
            if (response.ok) {
                onClose();
            } else {
                console.log("Échec de la modification du profil")
            }

        }catch (error){
            console.error("Erreur lors de la modification du profil :", error);
        }
    }
    if (!show) {
        return null;
    }
  return (
    <div className={`modal ${show ? 'show' : ''}`}>
      <div className="modal-content">
        <span className="close" onClick={onClose}>&times;</span>
        <h2>Modifier mon mot de passe</h2>
        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <label htmlFor="lastname">Mot de passe</label>
            <input type="password" id="password" name="password" value={formData.password} onChange={handleChange} />
          </div>
          <button type="submit">Valider</button>
        </form>
      </div>
    </div>
  );
};

export default ModalUpdatePassword;
