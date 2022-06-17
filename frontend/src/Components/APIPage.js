import { React, useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import {
  Box,
  Text,
  Spinner,
  Center,
  Flex,
  VStack,
  Button,
} from "@chakra-ui/react";
import axios from "axios";
import APIAccordion from "./APIAccordion";

function APIPage() {
  const { id } = useParams();

  const [APIData, setAPIData] = useState({});
  const [loading, setLoading] = useState(true);
  const [showButton, setShowButton] = useState(false);

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
      let remainder = count % 4;

      return (
        <APIAccordion
          key={item.request}
          id={id}
          item={item}
          index={count}
          remainder={remainder}
        />
      );
    });
  }

  useEffect(() => {
    setTimeout(() => {
      setShowButton(true);
    }, 10000);
  }, []);

  if (loading) {
    return (
      <Center>
        <Flex h={"92vh"} alignItems={"center"}>
          <VStack mx={50}>
            <Spinner size={"xl"} />
            {showButton && (
              <>
                <Text textAlign={"center"} pt={20} color={"red"}>
                  This is taking longer than expected...
                </Text>
                <Button
                  onClick={() => {
                    window.location.reload(false);
                  }}
                >
                  <Text textAlign={"center"}>Click here to refresh the page</Text>
                </Button>
                <Text textAlign={"center"} fontSize={"xs"} as={"i"}>
                  If the error persists, please reach out to one of our admins
                  at <b>testing@blabla.com</b> .
                </Text>
              </>
            )}
          </VStack>
        </Flex>
      </Center>
    );
  } else {
    return (
      <Box minH={"92vh"} my={50} mx={[50, 100, 150]}>
        <Text fontSize={["1.5rem", "2rem", "3rem", "4rem"]} color={"#222222"}>
          {APIData.title}
        </Text>

        <Text mt={5} ml={5} mb={"80px"} fontSize={["1rem", "1.25rem", "1.5rem"]} color={"gray"}>
          {APIData.description}
        </Text>

        {accordions}
      </Box>
    );
  }
}

export default APIPage;
