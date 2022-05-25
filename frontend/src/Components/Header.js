import { Box, Flex, Text, Heading, VStack } from "@chakra-ui/react";
import React from "react";

const Header = () => {
  return (
    <>
      <Box bg={"gray.300"} h={"92vh"}>
        <Flex h={"80%"} alignItems={"center"} justifyContent={"center"}>
          <VStack spacing={24}>
            <Heading>
              The Modern
              <br /> API Platform
            </Heading>
            <Text>Find, connect to, & manage thousands of APIs</Text>
          </VStack>
        </Flex>
      </Box>
    </>
  );
};

export default Header;
