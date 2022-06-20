import { React, useState } from "react";
import axios from "axios";
import {
  Box,
  Spinner,
  Text,
  Accordion,
  AccordionIcon,
  AccordionButton,
  AccordionItem,
  AccordionPanel,
  Center,
  Image,
} from "@chakra-ui/react";

function APIAccordion({ item, index, id, remainder }) {
  const [loading, setLoading] = useState(true);
  const [response, setResponse] = useState("");
  const [isOpen, setIsOpen] = useState(false);

  const fetchData = async () => {
    if (!isOpen) {
      const res = await axios.get(
        `https://papi-project.herokuapp.com/api/v1/apidata/${id}/${index}`
      );
      setResponse(res.data);
      setLoading(false);
      setIsOpen(true);
    }
  };

  let accordionBg;

  if (remainder === 0) {
    accordionBg = "gray.200";
  } else {
    accordionBg = "white";
  }

  function handleJSON(jsonObject, indent, key) {
    let entries = Object.entries(jsonObject);
    return entries.map((item, index) => {
      let newKey = `${key}${index}${item[0]}`;
      let val = typeof item[1];
      switch (val) {
        case "object":
          if (Array.isArray(item[1])) {
            return (
              <>
                <Text key={newKey} ml={indent} fontWeight={"bold"}>
                  {item[0]}:
                </Text>
                {handleArray(item[1], indent + 4, newKey)}
              </>
            );
          } else {
            return (
              <>
                <Text key={newKey} ml={indent} fontWeight={"bold"}>
                  {item[0]}:
                </Text>
                {handleJSON(item[1], indent + 4, newKey)}
              </>
            );
          }
        default:
          if (itemIsImage(item[1])) {
            return (
              <>
                <Text key={newKey} ml={indent} fontWeight={"bold"}>
                  {item[0]}:
                </Text>
                <Image
                  mx={"auto"}
                  maxH={"50vh"}
                  key={`${newKey}itemimage`}
                  src={item[1]}
                />
                <Text textAlign={"center"} key={`${newKey}item`}>
                  {item[1]}
                </Text>
              </>
            );
          } else if (item[0] === "lastUpdate") {
            let date = new Date(item[1]);
            return (
              <>
                <Text key={newKey} ml={indent} fontWeight={"bold"}>
                  {item[0]}:
                </Text>
                <Text key={`${newKey}item`} ml={indent + 4}>
                  {date.toLocaleDateString()}
                </Text>
              </>
            );
          } else {
            return (
              <>
                <Text key={newKey} ml={indent} fontWeight={"bold"}>
                  {item[0]}:
                </Text>
                <Text key={`${newKey}item`} ml={indent + 4}>
                  {item[1]}
                </Text>
              </>
            );
          }
      }
    });
  }

  function itemIsImage(item) {
    let imageSuffixs = ["png", "jpg", "gif"];
    let hasImg = false;
    imageSuffixs.forEach((suf) => {
      if (JSON.stringify(item).endsWith(`${suf}"`)) {
        hasImg = true;
      }
    });
    return hasImg;
  }

  function handleArray(arrayObject, indent, key) {
    return arrayObject.map((item, index) => {
      let val = typeof item;
      let newKey = `${key}${index}`;
      switch (val) {
        case "object":
          if (Array.isArray(item)) {
            return handleArray(item, indent + 4, newKey);
          } else {
            return handleJSON(item, indent + 4, newKey);
          }
        default:
          if (itemIsImage(item)) {
            return (
              <>
                <Image key={`${newKey}image`} src={item} />
                <Text key={newKey}>{item}</Text>
              </>
            );
          } else {
            return (
              <>
                <Text key={newKey} ml={indent}>
                  {item}
                </Text>
              </>
            );
          }
      }
    });
  }

  return (
    <>
      <Box bg={accordionBg}>
        <Accordion allowToggle mb={"5"}>
          <AccordionItem>
            <AccordionButton onClick={fetchData}>
              <Box flex="1" textAlign="left">
                <Text fontSize={["1rem", "1.25rem"]} color={"#222222"}>
                  {item.request}
                </Text>
              </Box>
              <AccordionIcon />
            </AccordionButton>
            <AccordionPanel pb={4} onClick={fetchData} p={1}>
              <Box bg={"white"} p={5}>
                {loading && (
                  <>
                    <Center my={20}>
                      <Spinner size={"xl"} />
                    </Center>
                  </>
                )}

                {!loading && handleJSON(response, 0, "")}
              </Box>
            </AccordionPanel>
          </AccordionItem>
        </Accordion>
      </Box>
    </>
  );
}

export default APIAccordion;
