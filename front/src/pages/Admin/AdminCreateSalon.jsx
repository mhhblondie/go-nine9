import React, {useState} from 'react'
import { getCookie, setCookie } from '../../AuthContext/AuthContext';
import { useNavigate } from 'react-router-dom';


const CreateSalonAdmin = () => {
    const [salon, setSalon] = useState({
        Name: '',
        Address: '',
        Phone: '',
        Description: '',
        user : []
    });
    const [errors, setErrors] = useState("")

    const token = getCookie("authToken")
    const navigate = useNavigate();



    // Change les valeurs du salon
    const handleInputChange = (event) => {
        const { name, value } = event.target;
        setSalon({ ...salon, [name]: value });
      };

    const NewSalon = async() =>{
        const {Name, Address, Phone, user, Description} = salon;
        const requestData = {Name, Address, Phone, user, Description};
        try {
            const response = await fetch('http://localhost:8097/api/admin/salons', {
              method: 'POST',
              headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`,
              },
              mode: 'cors',
              body: JSON.stringify(requestData),
            });
            
        
            const responseData = await response.json(); // Stocker les données de réponse dans une variable

            if (!response.ok) {
                console.log(responseData); 
                setErrors(responseData.message || 'Échec de la requête d\'inscription');
            } else {
                console.log(responseData); 
                setErrors("");
                navigate("/administrateur/AdminDashboard");
            }
          } catch (err) {
            setErrors(err.message || 'Une erreur inattendue s\'est produite');
          }

    }

    // Change les valeurs d'une personne du staff
    const handleStaffInputChange = (index, event) => {
        const { name, value } = event.target;
        const updatedStaff = [...salon.user];
        updatedStaff[index][name] = value;
        setSalon({ ...salon, user: updatedStaff });
      };

      //Ajoute une nouvelle entrée de input
      const addStaff = () => {
        setSalon({
          ...salon,
          user: [
            ...salon.user,
            { Lastname: '', Firstname: '', Email: '', Password: '' , Roles : "staff"},
          ],
        });
      };

      //retire un staff
      const removeStaff = (index) => {
        const updatedStaff = [...salon.user];
        updatedStaff.splice(index, 1);
        setSalon({ ...salon, user: updatedStaff });
      };
    

      
  return (
    <div>
        <h1> Créer un salon </h1>
        {errors.length > 0 && <h3>{errors}</h3>}
        <div>
            <label>Nom du salon</label>
            <input name="Name" type="text" placeholder='Ex: Mon salon' value={salon.Name} onChange={handleInputChange}/>
            <label>Adresse</label>
            <input name="Address" type="text" placeholder='Ex: 10 rue de Louvois 75002 Paris' value={salon.Address} onChange={handleInputChange}/>
            <label>Numéro de téléphone</label>
            <input name="Phone" type="tel" placeholder="Ex: 667" value={salon.Phone} onChange={handleInputChange}/>
            <label>Description</label>
            <input name="Description" type="textarea" placeholder="Ex: Description..." onChange={handleInputChange}/>
            <button onClick={addStaff}>Ajouter des collaborateurs</button>
            {salon.user.map((staffMember, index) => (
          <div key={index}>
          
          <h3>Staff {index+1}</h3>
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
            
            <button onClick={NewSalon}>Créer ton nouveau salon</button>
            </div>
      
    </div>
  )
}

export default CreateSalonAdmin
