import { Box, TextField, Button } from "@mui/material";
import React from "react";
import { useState } from "react";
import axios from "axios";

const TesDNA = () => {
  const [isHasil, setIsHasil] = useState(false);
  const [hasil, setHasil] = useState();
  const [fileChoosen, setFileChoosen] = useState(false);
  const [currFile, setCurrFile] = useState();

  const fileInputPengguna = document.getElementById("fileInputPengguna");
  const namaPengguna = document.getElementById("namaPengguna");
  const prediksiPenyakit = document.getElementById("prediksiPenyakit");

  const buttonSubmitHandler = () => {
    // TODO : submit data to API
    try{
      var f = fileInputPengguna.files[0];
      if (f && namaPengguna.value && prediksiPenyakit.value) {
        var r = new FileReader();
        r.onload = function (e) {
          var contents = e.target.result;
          console.log(contents);
          const tesData = async () => {
            const { data: hasilTes } = await axios.post(
              `http://localhost:8080/tesDNA/${namaPengguna.value}/${prediksiPenyakit.value}/${contents}`
            );
            setHasil(hasilTes);
            console.log(hasilTes);
            setIsHasil(true);
          };
          tesData();
          alert("Test berhasil");
        };
        r.readAsText(f);
      } else {
        alert("Isi semua data terlebih dahulu");
      }
    }catch(e){
      alert("Isi semua data terlebih dahulu");
    }
  };

  const changeFileInputHandler = (e) => {
    if (e.target.files[0]) {
      setFileChoosen(true);
      setCurrFile(e.target.files[0].name);
    } else {
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
          <h2>Tes DNA</h2>
        </Box>
        <Box
          sx={{
            display: "flex",
            justifyContent: "center",
            marginBottom: "1rem",
          }}
        >
          <TextField
            id="namaPengguna"
            label="Nama Pengguna"
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
            <input
              type="file"
              id="fileInputPengguna"
              hidden
              onChange={changeFileInputHandler}
            />
          </Button>
        </Box>
        <Box
          sx={{
            display: "flex",
            justifyContent: "center",
            marginBottom: "1rem",
          }}
        >
          <TextField
            id="prediksiPenyakit"
            label="Prediksi Penyakit"
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
          }}
        >
          <Button
            label="Submit"
            variant="outlined"
            onClick={buttonSubmitHandler}
          >
            Submit
          </Button>
        </Box>
        {isHasil ? (
          <Box
            sx={{
              display: "flex",
              justifyContent: "center",
              border: "1px black",
              borderRadius: "10px",
              backgroundColor: "white",
              width: "100%",
              marginTop: "1rem",
            }}
          >
            <h3>
              {hasil.Tanggal} - {hasil.Nama} - {hasil.Penyakit} -{" "}
              {hasil.diagnosis} - {hasil.persentase}%
            </h3>
          </Box>
        ) : null}
      </Box>
    </div>
  );
};

export default TesDNA;
