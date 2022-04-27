import "../styles/Home.css";
import { Box, TextField } from "@mui/material";
import { useState, useEffect } from "react";

const Home = () => {

  const [datas, setDatas] = useState([]);
  // const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      const {data: hasil} = await axios.get(`localhost:8080/diagnosis/all`);
      setDatas(hasil.Hasil);
      
    };
    fetchData();
  }, []);

  // TODO : get data from API
  const getDataHandler = (e) => {
    const searchKeyword = e.target.value;

    if(searchKeyword){
      // TODO : get data from API with search keyword
      const fetchData = async () => {
        const {data: hasil} = await axios.get(`localhost:8080/diagnosis/${searchKeyword}`);
        setDatas(hasil.Hasil);
      };
      fetchData();
    }else{
      // TODO : get data from API
      const fetchData = async () => {
        const {data: hasil} = await axios.get(`localhost:8080/diagnosis/all`);
        setDatas(hasil.Hasil);
        
      };
      fetchData();
    }
  }

  // const datas = [
  //   {
  //     nama: "Aldwin",
  //     tanggal: "2020-01-01",
  //     penyakit: "Penyakit 1",
  //     hasil: "Positif",
  //   }
  // ]


  return (
    <div className="main">
      <h1>DNA Sequence Matcher</h1>
      <Box
        component="div"
        sx={{
          display: "flex",
          justifyContent: "center",
          //   alignItems: "left",
          flexDirection: "column",
          border: "3px black",
          borderRadius: "10px",
          height: "auto",
          width: "50%",
          padding: "1rem",
          backgroundColor: "lightblue",
        }}
      >
        <Box
          sx={{
            display: "flex",
            flexDirection: "row",
            justifyContent: "center",
          }}
        >
          <h2>Cari Hasil Tes DNA</h2>
        </Box>
        <Box
          sx={{
            display: "flex",
            justifyContent: "center",
            marginBottom: "1rem",
          }}
        >
          <TextField
            id="search"
            label="Tanggal/Nama"
            variant="outlined"
            sx={{
              width: "50%",
              margin: "1rem",
            }}
            onChange={getDataHandler}// getData
          />
        </Box>
        <Box
          sx={{
            display: "flex",
            justifyContent: "center",
          }}
        >
          {datas.map((data) => (
            <Box
              sx={{
                display: "flex",
                justifyContent: "center",
                border: "1px black",
                borderRadius: "10px",
                backgroundColor: "white",
                width: "100%",
              }}
            >
              <h3>
                {data.tanggal} - {data.nama} - {data.penyakit} - {data.diagnosis}
              </h3>
            </Box>
          ))}
        </Box>
      </Box>
    </div>
  );
};

export default Home;
