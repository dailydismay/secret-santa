import { styled } from "@nextui-org/react";

export const Layout = styled("div", {
  display: "flex",
  flexDirection: "column",
  alignItems: "center",
  minHeight: "100vh",
  margin: 0,
});

export const Content = styled("main", {
  "@xsMin": {
    flex: 1,
    height: "100%",
    maxWidth: "1400px",
    width: "100%",
  },
});
