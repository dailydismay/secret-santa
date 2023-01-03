import { Navbar, styled } from "@nextui-org/react";
import { useStore } from "effector-react";
import { useLocation } from "react-router-dom";
import { $translation } from "store/i18n";
import { Visitor } from "store/visitor/types";
import { loggedInNavItems } from "./items";
import { NavItemLink } from "./nav-item-link";
import { Profile } from "./profile";

export type NavBarContentProps = {
  isLoggedIn: boolean;
  visitor: Visitor | null;
};

const ContentWrapper = styled("div", {
  "@xsMax": {
    display: "none",
  },
});

const ProfileWrapper = styled("div", {
  display: "none",
  "@xsMin": {
    display: "block",
    marginLeft: "$5",
    marginRight: "$5",
  },
});

export const VisitorNavBarContent: React.FC<{ visitor: Visitor }> = ({
  visitor,
}) => {
  const location = useLocation();
  const $t = useStore($translation);

  return (
    <>
      {loggedInNavItems.map(({ translationItem, path }, key) => (
        <Navbar.Item hideIn={"xs"} key={key}>
          <NavItemLink
            path={path}
            text={$t[translationItem]}
            currentPath={location.pathname}
          ></NavItemLink>
        </Navbar.Item>
      ))}
      <Navbar.Item
        css={{ cursor: "pointer" }}
        hideIn={"xs"}
        isActive={location.pathname === "/profile"}
      >
        <ProfileWrapper>
          <Profile visitor={visitor}></Profile>
        </ProfileWrapper>
      </Navbar.Item>
    </>
  );
};

export const NavBarContent: React.FC<NavBarContentProps> = ({
  isLoggedIn,
  visitor,
}) => (
  <Navbar.Content activeColor={"secondary"} variant={"underline-rounded"}>
    <Navbar.Brand showIn={"xs"}>
      <Navbar.Toggle aria-label="toggle navigation" />
    </Navbar.Brand>
    {isLoggedIn && visitor && (
      <VisitorNavBarContent visitor={visitor}></VisitorNavBarContent>
    )}
  </Navbar.Content>
);
