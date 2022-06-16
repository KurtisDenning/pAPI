import React from "react";
import { Box, Flex, Heading } from "@chakra-ui/react";

const Footer = () => {
  return (
    <Box bg={"#222222"}>
      <Flex h={"33vh"} alignItems={"center"} justifyContent={"center"}>
        <Heading color={"white"} fontSize={["2rem", "3rem"]} >
          pAPI
        </Heading>
      </Flex>
    </Box>
  );
};

export default Footer;
