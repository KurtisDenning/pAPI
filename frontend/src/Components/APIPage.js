import { React, useEffect } from "react";
import { useParams } from "react-router-dom"
import { Heading, Box, Text } from "@chakra-ui/react";
import axios from "axios";

function APIPage() {
  const { id } = useParams();
  // let APIData = [];

  // useEffect(() => {
  //   const fetchAPIData = async () => {
  //     const res = await axios.get(
  //       `https://papi-project.herokuapp.com/api/v1/apidata/${id}`
  //     );
  //     APIData = res.data;
  //   };
  //   fetchAPIData();
  // }, [])

  // const API = APIData.map((item) => {
  //   console.log(item.description);
  // });

  return (
    <Box my={50} mx={[50, 100, 150]}>
      <Heading as={"h3"} size={"lg"}>Title</Heading>
      <Text mt={4} mb={"80px"}>
        Description
      </Text>
    </Box>
  );
}

export default APIPage;
