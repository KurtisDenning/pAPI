import React from "react";
import APICard from "./APICards";
import APIData from "../Dummy-Json/API-data";
import { Center, Heading, Box, Flex } from "@chakra-ui/react";

const apiCards = APIData.map((item) => <APICard key={item.id} {...item} />);

const APIs = () => {
  return (
    <Box m={50}>
      <Center mb={50}>
        <Heading as={"h3"} size={"lg"}>
          API's {/*need to set state to for what results are being filtered*/}
        </Heading>
      </Center>
      {/*Need to add a search bar.*/}
      {apiCards}
    </Box>
  );
};

export default APIs;
