import { Box, TextField, Button } from "@mui/material";

const TesDNA = () => {
  const buttonSubmitHandler = (e) => {
    e.preventDefault();

    // TODO : submit data to API
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
            id="nama-pengguna"
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
            {""}
            Upload a file
            <input type="file" hidden />
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
            id="prediksi-penyakit"
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
          <Button label="Submit" variant="outlined">
            Submit
          </Button>
        </Box>
      </Box>
    </div>
  );
};

export default TesDNA;
