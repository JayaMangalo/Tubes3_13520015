import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Layout from "./pages/Layout";
import Home from "./pages/Home";
import TambahPenyakit from "./pages/TambahPenyakit";
import TesDNA from "./pages/TesDNA";
import NoPage from "./pages/NoPage.js";
import "./styles/index.css";

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<Home />} />
          <Route path="tambahpenyakit" element={<TambahPenyakit />} />
          <Route path="tesdna" element={<TesDNA />} />
          <Route path="*" element={<NoPage />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(<App />);
