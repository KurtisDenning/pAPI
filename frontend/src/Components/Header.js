import { Box, Flex, Text, Heading, VStack, Input } from "@chakra-ui/react";
import React from "react";

const Header = () => {
  return (
    <>
      <Box h={"92vh"} mx={50}>
        <Flex h={"80%"} alignItems={"center"} justifyContent={"center"}>
          <VStack spacing={10}>
            <Heading as={"h1"} size={"3xl"}>
              The Modern
              <br /> API Platform
            </Heading>
            <Text>Find, connect to, & manage thousands of APIs</Text>
          </VStack>
        </Flex>
        <Flex justifyContent={"center"}>
          <Input
            w={"80vw"}
            isInvalid
            errorBorderColor="black"
            placeholder="Search for an API here"
          ></Input>
        </Flex>
      </Box>
    </>
  );
};

export default Header;
