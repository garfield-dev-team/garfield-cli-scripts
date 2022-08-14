import * as React from "react";
import s from "./style.module.less";

type IProps = React.PropsWithChildren<{}>;

const App: React.FC<IProps> = ({ children }) => {
    return <div></div>
}

export default React.memo(App);

