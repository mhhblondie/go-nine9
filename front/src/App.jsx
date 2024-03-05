import { useContext } from 'react';
import { Routes, Route, BrowserRouter } from 'react-router-dom';
import { AuthContext, AuthProvider } from './AuthContext/AuthContext';
import Login from './pages/login/Login';
import Home from './pages/home/Home';
import Dashboard from './pages/Manager/Dashboard';
import CreateSalons from './pages/Manager/CreateSalons';
import Register from './pages/register/Register';
import Salon from './pages/salons/Salon';
import AddStaff from './pages/Manager/AddStaff';
import ModifyMySalon from './pages/Manager/ModifySalon';
import Reservation from './pages/reservation/Reservation';
import Confirmation from './pages/confirmation/Confirmation';
import Navbar from './components/Navbar/Navbar';
import AdminDashboard from './pages/Admin/AdminDashboard';
import CreateSalonAdmin from './pages/Admin/AdminCreateSalon';
function App() {
  const { isAuthenticated, isManager } = useContext(AuthContext);

  const ManagerRoutes = (
    <Routes>
        <Route path="/admin/*" element={<StaffRoutes />} />
      <Route path="/*" element={<Home />} />
    </Routes>
  );

  const publicRoutes = (
    <Routes>
      <Route path="/login" element={<Login />} />
      <Route path="/register" element={<Register />} />
      <Route path="/" element={<Home />} />
      <Route path="/salons/:id" element={<Salon />} />
      <Route path="salons/:id/reservation" element={<Reservation />} />
      <Route path="/confirmation" element={<Confirmation />} />
      <Route path="/administrateur" element={<Confirmation />} />
      <Route path="/administrateur/AdminDashboard" element={<AdminDashboard />} />
      <Route path="/administrateur/createSalon" element={<CreateSalonAdmin />} />
      <Route path="/administrateur/ModifySalon" element={<Confirmation />} />
    </Routes>
  );

  const Layout = () => {
    return (
      <div>
        <BrowserRouter>
          <AuthProvider>
          <Navbar/>
            {isManager ? ManagerRoutes : isAuthenticated ? publicRoutes : publicRoutes}
          </AuthProvider>
        </BrowserRouter>
      </div>
    );
  };

  return <Layout />;
}

function StaffRoutes() {
  return (
    <Routes>
      <Route path="/dashboard" element={<Dashboard />} />
      <Route path="/create" element={<CreateSalons />} />
      <Route path="/addStaff" element={<AddStaff />} />
      <Route path="/modify" element={<ModifyMySalon />} />
    </Routes>
  );
}

function AdminRoutes() {
  return (
    <Routes>
      <Route path="/" element={<Dashboard />} />
      <Route path="/create" element={<CreateSalons />} />
      <Route path="/addStaff" element={<AddStaff />} />
      <Route path="/modify" element={<ModifyMySalon />} />
    </Routes>
  );
}

export default App;
