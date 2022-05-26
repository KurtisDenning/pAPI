import React from "react";
import { Box, Flex, Text } from "@chakra-ui/react";

const Footer = () => {
  return (
    <Box bg={"#222222"}>
      <Flex h={"33vh"} alignItems={"center"} justifyContent={"center"}>
        <Text color={"white"} fontSize={"3xl"}>
          pAPI
        </Text>
      </Flex>
    </Box>
  );
};

export default Footer;
