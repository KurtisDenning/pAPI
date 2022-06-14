import { React, useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { Heading, Box, Text, Spinner } from "@chakra-ui/react";
import axios from "axios";
import APIAccordion from "./APIAccordion";

function APIPage() {
  const { id } = useParams();

  const [APIData, setAPIData] = useState({});
  const [loading, setLoading] = useState(true);
  let accordions;

  useEffect(() => {
    const fetchAPIData = async () => {
      const res = await axios.get(
        `https://papi-project.herokuapp.com/api/v1/apidata/${id}`
      );
      setAPIData(res.data);
      setLoading(false);
    };
    fetchAPIData();
  }, [id]);

  let requests = APIData.requests;

  if (!loading) {
    let count = -1;
    accordions = requests.map((item) => {
      count++;
      return (
        <APIAccordion key={item.request} id={id} item={item} index={count} />
      );
    });
  }

  if (loading) {
    return (
      <Box my={50} mx={[50, 100, 150]}>
        <Spinner size={"xl"} />
      </Box>
    );
  } else {
    return (
      <Box my={50} mx={[50, 100, 150]}>
        <Heading as={"h3"} size={"lg"}>
          {APIData.title}
        </Heading>

        <Text mt={4} mb={"80px"}>
          {APIData.description}
        </Text>

        {accordions}
      </Box>
    );
  }
}

export default APIPage;
