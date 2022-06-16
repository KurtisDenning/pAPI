import React from "react";
import { Box, Flex, Heading } from "@chakra-ui/react";

const Footer = () => {
  return (
    <Box bg={"#222222"}>
      <Flex h={"33vh"} alignItems={"center"} justifyContent={"center"}>
        <Heading color={"white"} as={"h2"} size={"2xl"}>
          pAPI
        </Heading>
      </Flex>
    </Box>
  );
};

export default Footer;
