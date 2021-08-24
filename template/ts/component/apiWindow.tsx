import React from "react";

const onClickTest = () => {
    alert("The button has been pushed.");
}

const ceilingStatus = (domain: string, status: string) => {
    fetch(`${domain}/maid/scenario/ceiling/${status}`)
        .then();
}

type rowParam = {
    onClick: any;
    url: string;
    description: string;
}

const generateRow = (param: rowParam): JSX.Element => {
    return (
        <tr>
            <td>
                <button
                    className="mx-auto btn btn-primary"
                    onClick={() => { param.onClick() }}
                >
                    {"Execute"}
                </button>
            </td>
            <td>
                {param.url}
            </td>
            <td>
                {param.description}
            </td>
        </tr>
    )
}

type ApiWindowProps = {
    domain: string;
}

export const ApiWindow: React.FC<ApiWindowProps> = (props: ApiWindowProps) => {
    return (
        <div>
            <div className="w-75 mx-auto">
                {"div"}
            </div>
            <table className="w-75 table table-bordered border border-dark mx-auto">
                <thead>
                    <tr>
                        <th></th>
                        <th>URL</th>
                        <th>Description</th>
                    </tr>
                </thead>
                <tbody>
                    {
                        generateRow(
                            {
                                onClick: () => { ceilingStatus(props.domain, 'on') },
                                description: 'Turn on ceiling light.',
                                url: '/maid/scenario/ceiling/on',
                            })
                    }
                    {
                        generateRow(
                            {
                                onClick: () => { ceilingStatus(props.domain, 'off') },
                                description: 'Turn off ceiling light.',
                                url: '/maid/scenario/ceiling/off',
                            })
                    }
                    {
                        generateRow(
                            {
                                onClick: () => { fetch(`${props.domain}/maid/light/ip`).then(); },
                                description: 'Refresh IP of lightbulbs.',
                                url: '/maid/light/ip',
                            })
                    }
                </tbody>
            </table>
        </div>
    );
}