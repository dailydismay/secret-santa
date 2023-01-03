import React from "react";
import { Link as RouterLink } from "react-router-dom";
import { useEvent, useStore } from "effector-react";
import {
  Avatar,
  Container,
  Divider,
  Link,
  Navbar,
  Text,
} from "@nextui-org/react";
import { $isLoggedIn, $visitor } from "store/visitor";
import { $translation, TranslationItem } from "store/i18n";
import { NavBarContent } from "./content";
import { $logoutClicked } from "store/visitor/events";

const Collapse: React.FC = () => {
  const $t = useStore($translation);
  const visitor = useStore($visitor);
  const logoutClicked = useEvent($logoutClicked);

  return (
    <Navbar.Collapse
      css={
        {
          // "& ul": {
          // height: "100%",
          // },
        }
      }
    >
      <Navbar.CollapseItem>
        <Container>
          <Link
            color="inherit"
            css={{
              paddingY: "$5",
              minWidth: "100%",
              fontSize: "$xl",
            }}
            href="#"
          >
            {$t[TranslationItem.NavGroups]}
          </Link>
        </Container>
      </Navbar.CollapseItem>
      <Navbar.CollapseItem>
        {visitor && (
          <Container display="flex" justify={"space-between"}>
            <Link
              color="inherit"
              css={{
                paddingY: "$5",
                fontSize: "$xl",
                minWidth: "100%",
              }}
              href="#"
            >
              <RouterLink style={{ color: "black" }} to={"/profile"}>
                {$t[TranslationItem.NavProfile]}
              </RouterLink>
            </Link>
          </Container>
        )}
      </Navbar.CollapseItem>
      <Divider></Divider>
      <Navbar.CollapseItem>
        <Container
          onClick={() => logoutClicked()}
          css={{ marginTop: "$5", fontSize: "$xl" }}
        >
          {$t[TranslationItem.NavLogout]}
        </Container>
      </Navbar.CollapseItem>
    </Navbar.Collapse>
  );
};

export const Header: React.FC = () => {
  const $t = useStore($translation);
  const isLoggedIn = useStore($isLoggedIn);
  const visitor = useStore($visitor);

  const emojis = ["🎁", "🎅"];

  return (
    <>
      <Navbar
        disableBlur={false}
        disableShadow={false}
        shouldHideOnScroll
        variant="sticky"
      >
        <Navbar.Brand>
          <Text b css={{ fontSize: "$lg" }} color="inherit">
            {emojis[Math.floor(Math.random() * emojis.length)]}{" "}
            {$t[TranslationItem.NavBrand]}
          </Text>
        </Navbar.Brand>
        <NavBarContent
          isLoggedIn={isLoggedIn}
          visitor={visitor}
        ></NavBarContent>
        <Collapse />
      </Navbar>
    </>
  );
};
