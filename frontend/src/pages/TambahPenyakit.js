import { Button, Box, TextField } from "@mui/material";
import { useState } from "react";
import "../styles/TambahPenyakit.css";
import axios from 'axios'

const TambahPenyakit = () => {

  const [fileChoosen, setFileChoosen] = useState(false);
  const [currFile, setCurrFile] = useState();

  const namaPenyakit = document.getElementById("namaPenyakit");
  const fileInput = document.getElementById("fileInput");

  const buttonSubmitHandler = () => {

    // TODO : submit data to API
    console.log(fileInput.files[0]);
    var f = fileInput.files[0]
    if (f) {
      var r = new FileReader();
      r.onload = function(e) { 
          var contents = e.target.result;             
          console.log(contents)
          const insertData = async () => {
            await axios.post(`http://localhost:8080/inputPenyakit/${namaPenyakit.value}/${contents}`);
          };
          insertData();
          alert("Penyakit berhasil di upload");
      }
      r.readAsText(f);
    } else { 
      alert("Failed to load file");
    }
    
  };

  const changeFileInputHandler = (e) => {
    if(e.target.files[0]){
      setFileChoosen(true);
      setCurrFile(e.target.files[0].name);
    }else{
      setFileChoosen(false);
    }
  };

  return (
    <div className="main">
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
          <h2>Tambah Penyakit</h2>
        </Box>
        <Box
          sx={{
            display: "flex",
            justifyContent: "center",
            marginBottom: "1rem",
          }}
        >
          <TextField
            id="namaPenyakit"
            label="Nama Penyakit"
            variant="outlined"
            sx={{
              width: "50%",
              margin: "1rem",
            }}
          />
        </Box>
        <Box
          sx={{
            display: "flex",
            justifyContent: "center",
            flexDirection: "row",
            marginBottom: "2rem",
          }}
        >
          <label
            sx={{
              display: "flex",
              justifyContent: "flex-start",
            }}
          >
            Sequence DNA:
          </label>
          <Button
            variant="contained"
            component="label"
            color="primary"
            sx={{ marginLeft: "4rem" }}
          >
            {fileChoosen ? `${currFile}` : "Pilih File"}
            <input type="file" id="fileInput" hidden onChange={changeFileInputHandler}/>
          </Button>
        </Box>
        <Box
          sx={{
            display: "flex",
            justifyContent: "center",
            flexDirection: "row",
          }}
        >
          <Button label="Submit" variant="outlined" onClick={buttonSubmitHandler}>
            Submit
          </Button>
        </Box>
      </Box>
    </div>
  );
};

export default TambahPenyakit;
