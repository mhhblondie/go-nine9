import React, { useState } from 'react';
import { getCookie } from '../../AuthContext/AuthContext';
import { Link, useNavigate } from "react-router-dom";

const AddStaff = () => {
    const [staff, setStaff] = useState([
        { Lastname: '', Firstname: '', Email: '', Password: '', Roles: "staff" }
    ]);

    const [errors, setErrors] = useState("");

    const token = getCookie("authToken");
    const navigate = useNavigate();

    // Change les valeurs d'une personne du staff
    const handleStaffInputChange = (index, event) => {
        const { name, value } = event.target;
        const updatedStaff = [...staff];
        updatedStaff[index][name] = value;
        setStaff(updatedStaff);
    };

    // Ajoute une nouvelle entrée de input
    const addStaff = () => {
        setStaff([
            ...staff,
            { Lastname: '', Firstname: '', Email: '', Password: '', Roles: "staff" },
        ]);
    };

    // Retire un staff
    const removeStaff = (index) => {
        const updatedStaff = [...staff];
        updatedStaff.splice(index, 1);
        setStaff(updatedStaff);
    };

    const handleAddStaff = async() =>{
        try {
            const response = await fetch('http://localhost:8097/api/management/salons/staff', {
              method: 'POST',
              headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`,
              },
              mode: 'cors',
              body: JSON.stringify(staff),
            });
        
            if (!response.ok) {
              const errorResponse = await response.json();
              console.log(errorResponse)
              setErrors(errorResponse.message || 'Échec de la requête d\'inscription');
            }
            
            if(response.ok){
                setErrors("")
                navigate("/admin/")
            }
           
            

          } catch (err) {
            setErrors(err.message || 'Une erreur inattendue s\'est produite');
          }

    }

    return (
        <div>
            <h1>Ajouter des employés</h1>
            <button onClick={addStaff}>Ajouter des collaborateurs</button>
            <button onClick={handleAddStaff}>Enregistrer</button>
            {staff.map((staffMember, index) => (
                <div key={index}>
                    <h3>Staff {index + 1}</h3>
                    <label>Nom</label>
                    <input
                        name="Lastname"
                        type="text"
                        value={staffMember.Lastname}
                        onChange={(e) => handleStaffInputChange(index, e)}
                    />
                    <label>Prénom</label>
                    <input
                        name="Firstname"
                        type="text"
                        value={staffMember.Firstname}
                        onChange={(e) => handleStaffInputChange(index, e)}
                    />
                    <label>Email</label>
                    <input
                        name="Email"
                        type="email"
                        value={staffMember.Email}
                        onChange={(e) => handleStaffInputChange(index, e)}
                    />
                    <label>Mot de passe</label>
                    <input
                        name="Password"
                        type="password"
                        value={staffMember.Password}
                        onChange={(e) => handleStaffInputChange(index, e)}
                    />
                    <button onClick={() => removeStaff(index)}>Supprimer</button>
                </div>
            ))}
        </div>
    );
};

export default AddStaff;
